package net

import (
	"testing"
	"time"
)

func TestTCPListener_ListenAndAccept(t *testing.T) {
	n := NewNet()

	// Listen on an ephemeral port
	addr, err := n.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveTCPAddr() error = %v", err)
	}

	listener, err := n.ListenTCP("tcp", addr)
	if err != nil {
		t.Fatalf("ListenTCP() error = %v", err)
	}
	defer listener.Close()

	if listener == nil {
		t.Fatal("ListenTCP() returned nil")
	}

	// Get the actual address
	listenAddr := listener.Addr()
	if listenAddr == nil {
		t.Fatal("Addr() returned nil")
	}

	// Test Accept in a goroutine
	acceptDone := make(chan bool)
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Logf("Accept() error (may be expected after close): %v", err)
			return
		}
		conn.Close()
		acceptDone <- true
	}()

	// Give Accept time to start, then close the listener
	time.Sleep(10 * time.Millisecond)
	listener.Close()
}

func TestTCPListener_AcceptTCP(t *testing.T) {
	n := NewNet()

	// Listen on an ephemeral port
	addr, err := n.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveTCPAddr() error = %v", err)
	}

	listener, err := n.ListenTCP("tcp", addr)
	if err != nil {
		t.Fatalf("ListenTCP() error = %v", err)
	}
	defer listener.Close()

	listenAddr := listener.Addr().String()

	// Connect from a client in a goroutine
	go func() {
		time.Sleep(10 * time.Millisecond)
		conn, err := n.Dial("tcp", listenAddr)
		if err != nil {
			return
		}
		defer conn.Close()
		time.Sleep(100 * time.Millisecond)
	}()

	// Accept the connection
	conn, err := listener.AcceptTCP()
	if err != nil {
		t.Errorf("AcceptTCP() error = %v", err)
		return
	}
	defer conn.Close()

	if conn == nil {
		t.Fatal("AcceptTCP() returned nil")
	}
}

func TestTCPListener_SetDeadline(t *testing.T) {
	n := NewNet()

	// Listen on an ephemeral port
	addr, err := n.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveTCPAddr() error = %v", err)
	}

	listener, err := n.ListenTCP("tcp", addr)
	if err != nil {
		t.Fatalf("ListenTCP() error = %v", err)
	}
	defer listener.Close()

	// Set a deadline in the past to make Accept fail immediately
	err = listener.SetDeadline(time.Now().Add(-1 * time.Second))
	if err != nil {
		t.Errorf("SetDeadline() error = %v", err)
	}

	// Accept should fail with timeout
	_, err = listener.Accept()
	if err == nil {
		t.Error("Accept() should have failed with timeout, got nil error")
	}
}

func TestTCPListener_File(t *testing.T) {
	n := NewNet()

	// Listen on an ephemeral port
	addr, err := n.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveTCPAddr() error = %v", err)
	}

	listener, err := n.ListenTCP("tcp", addr)
	if err != nil {
		t.Fatalf("ListenTCP() error = %v", err)
	}
	defer listener.Close()

	// Get the underlying file
	file, err := listener.File()
	if err != nil {
		t.Errorf("File() error = %v", err)
		return
	}
	defer file.Close()

	if file == nil {
		t.Error("File() returned nil")
	}
}

func TestTCPListener_SyscallConn(t *testing.T) {
	n := NewNet()

	// Listen on an ephemeral port
	addr, err := n.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveTCPAddr() error = %v", err)
	}

	listener, err := n.ListenTCP("tcp", addr)
	if err != nil {
		t.Fatalf("ListenTCP() error = %v", err)
	}
	defer listener.Close()

	// Get the syscall connection
	syscallConn, err := listener.SyscallConn()
	if err != nil {
		t.Errorf("SyscallConn() error = %v", err)
		return
	}

	if syscallConn == nil {
		t.Error("SyscallConn() returned nil")
	}
}
