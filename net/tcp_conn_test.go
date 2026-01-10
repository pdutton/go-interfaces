package net

import (
	"io"
	"testing"
	"time"
)

func TestTCPConn_DialAndClose(t *testing.T) {
	n := NewNet()

	// Start a test TCP listener
	listener, err := n.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Dial the listener
	conn, err := n.DialTCP("tcp", nil, parseTCPAddr(t, addr))
	if err != nil {
		t.Fatalf("DialTCP() error = %v", err)
	}
	defer conn.Close()

	if conn == nil {
		t.Fatal("DialTCP() returned nil connection")
	}

	// Test Close
	err = conn.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
	}
}

func TestTCPConn_ReadWrite(t *testing.T) {
	n := NewNet()

	// Start a TCP listener
	listener, err := n.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Accept connections in a goroutine
	done := make(chan bool)
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Errorf("Accept() error = %v", err)
			return
		}
		defer conn.Close()

		// Read from client
		buf := make([]byte, 100)
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			t.Errorf("Read() error = %v", err)
			return
		}

		// Echo back
		_, err = conn.Write(buf[:n])
		if err != nil {
			t.Errorf("Write() error = %v", err)
		}

		done <- true
	}()

	// Dial the listener
	conn, err := n.DialTCP("tcp", nil, parseTCPAddr(t, addr))
	if err != nil {
		t.Fatalf("DialTCP() error = %v", err)
	}
	defer conn.Close()

	// Write to server
	message := []byte("hello")
	n1, err := conn.Write(message)
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	if n1 != len(message) {
		t.Errorf("Write() wrote %d bytes, want %d", n1, len(message))
	}

	// Read response
	buf := make([]byte, 100)
	n2, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		t.Fatalf("Read() error = %v", err)
	}

	if string(buf[:n2]) != "hello" {
		t.Errorf("Read() = %q, want %q", string(buf[:n2]), "hello")
	}

	<-done
}

func TestTCPConn_LocalRemoteAddr(t *testing.T) {
	n := NewNet()

	// Start a TCP listener
	listener, err := n.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Dial the listener
	conn, err := n.DialTCP("tcp", nil, parseTCPAddr(t, addr))
	if err != nil {
		t.Fatalf("DialTCP() error = %v", err)
	}
	defer conn.Close()

	// Test LocalAddr
	localAddr := conn.LocalAddr()
	if localAddr == nil {
		t.Error("LocalAddr() returned nil")
	}

	// Test RemoteAddr
	remoteAddr := conn.RemoteAddr()
	if remoteAddr == nil {
		t.Error("RemoteAddr() returned nil")
	}
}

func TestTCPConn_Deadlines(t *testing.T) {
	n := NewNet()

	// Start a TCP listener
	listener, err := n.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Dial the listener
	conn, err := n.DialTCP("tcp", nil, parseTCPAddr(t, addr))
	if err != nil {
		t.Fatalf("DialTCP() error = %v", err)
	}
	defer conn.Close()

	// Test SetDeadline
	deadline := time.Now().Add(1 * time.Second)
	err = conn.SetDeadline(deadline)
	if err != nil {
		t.Errorf("SetDeadline() error = %v", err)
	}

	// Test SetReadDeadline
	err = conn.SetReadDeadline(deadline)
	if err != nil {
		t.Errorf("SetReadDeadline() error = %v", err)
	}

	// Test SetWriteDeadline
	err = conn.SetWriteDeadline(deadline)
	if err != nil {
		t.Errorf("SetWriteDeadline() error = %v", err)
	}
}

func TestTCPConn_Options(t *testing.T) {
	n := NewNet()

	// Start a TCP listener
	listener, err := n.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Dial the listener
	conn, err := n.DialTCP("tcp", nil, parseTCPAddr(t, addr))
	if err != nil {
		t.Fatalf("DialTCP() error = %v", err)
	}
	defer conn.Close()

	// Test SetKeepAlive
	err = conn.SetKeepAlive(true)
	if err != nil {
		t.Errorf("SetKeepAlive() error = %v", err)
	}

	// Test SetKeepAlivePeriod
	err = conn.SetKeepAlivePeriod(30 * time.Second)
	if err != nil {
		t.Errorf("SetKeepAlivePeriod() error = %v", err)
	}

	// Test SetLinger
	err = conn.SetLinger(0)
	if err != nil {
		t.Errorf("SetLinger() error = %v", err)
	}

	// Test SetNoDelay
	err = conn.SetNoDelay(true)
	if err != nil {
		t.Errorf("SetNoDelay() error = %v", err)
	}

	// Test SetReadBuffer
	err = conn.SetReadBuffer(4096)
	if err != nil {
		t.Errorf("SetReadBuffer() error = %v", err)
	}

	// Test SetWriteBuffer
	err = conn.SetWriteBuffer(4096)
	if err != nil {
		t.Errorf("SetWriteBuffer() error = %v", err)
	}
}

func TestTCPConn_CloseReadWrite(t *testing.T) {
	n := NewNet()

	// Start a TCP listener
	listener, err := n.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Dial the listener
	conn, err := n.DialTCP("tcp", nil, parseTCPAddr(t, addr))
	if err != nil {
		t.Fatalf("DialTCP() error = %v", err)
	}
	defer conn.Close()

	// Test CloseWrite
	err = conn.CloseWrite()
	if err != nil {
		t.Errorf("CloseWrite() error = %v", err)
	}
}

// Helper function to parse TCP address
func parseTCPAddr(t *testing.T, addr string) *TCPAddr {
	t.Helper()
	n := NewNet()
	tcpAddr, err := n.ResolveTCPAddr("tcp", addr)
	if err != nil {
		t.Fatalf("ResolveTCPAddr() error = %v", err)
	}
	return tcpAddr
}
