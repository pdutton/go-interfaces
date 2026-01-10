package net

import (
	"testing"
	"time"
)

func TestUDPConn_DialAndListen(t *testing.T) {
	n := NewNet()

	// Listen for UDP packets
	addr, err := n.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveUDPAddr() error = %v", err)
	}

	listener, err := n.ListenUDP("udp", addr)
	if err != nil {
		t.Fatalf("ListenUDP() error = %v", err)
	}
	defer listener.Close()

	if listener == nil {
		t.Fatal("ListenUDP() returned nil")
	}

	listenAddr := listener.LocalAddr()
	if listenAddr == nil {
		t.Fatal("LocalAddr() returned nil")
	}

	// Dial the listener
	conn, err := n.DialUDP("udp", nil, listenAddr.(*UDPAddr))
	if err != nil {
		t.Fatalf("DialUDP() error = %v", err)
	}
	defer conn.Close()

	if conn == nil {
		t.Fatal("DialUDP() returned nil")
	}
}

func TestUDPConn_ReadWrite(t *testing.T) {
	n := NewNet()

	// Create a UDP listener
	addr, err := n.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveUDPAddr() error = %v", err)
	}

	listener, err := n.ListenUDP("udp", addr)
	if err != nil {
		t.Fatalf("ListenUDP() error = %v", err)
	}
	defer listener.Close()

	listenAddr := listener.LocalAddr().(*UDPAddr)

	// Create a client conn
	conn, err := n.DialUDP("udp", nil, listenAddr)
	if err != nil {
		t.Fatalf("DialUDP() error = %v", err)
	}
	defer conn.Close()

	// Write from client
	message := []byte("hello udp")
	nwrite, err := conn.Write(message)
	if err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	if nwrite != len(message) {
		t.Errorf("Write() wrote %d bytes, want %d", nwrite, len(message))
	}

	// Read on server
	buf := make([]byte, 100)
	listener.SetReadDeadline(time.Now().Add(1 * time.Second))
	nread, err := listener.Read(buf)
	if err != nil {
		t.Errorf("Read() error = %v", err)
		return
	}

	if string(buf[:nread]) != "hello udp" {
		t.Errorf("Read() = %q, want %q", string(buf[:nread]), "hello udp")
	}
}

func TestUDPConn_ReadFromWriteTo(t *testing.T) {
	n := NewNet()

	// Create a UDP listener
	addr, err := n.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveUDPAddr() error = %v", err)
	}

	listener, err := n.ListenUDP("udp", addr)
	if err != nil {
		t.Fatalf("ListenUDP() error = %v", err)
	}
	defer listener.Close()

	listenAddr := listener.LocalAddr().(*UDPAddr)

	// Create a client conn
	clientAddr, err := n.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveUDPAddr() error = %v", err)
	}

	client, err := n.ListenUDP("udp", clientAddr)
	if err != nil {
		t.Fatalf("ListenUDP() error = %v", err)
	}
	defer client.Close()

	// Write from client using WriteTo
	message := []byte("test message")
	nwrite, err := client.WriteTo(message, listenAddr)
	if err != nil {
		t.Fatalf("WriteTo() error = %v", err)
	}

	if nwrite != len(message) {
		t.Errorf("WriteTo() wrote %d bytes, want %d", nwrite, len(message))
	}

	// Read on server using ReadFrom
	buf := make([]byte, 100)
	listener.SetReadDeadline(time.Now().Add(1 * time.Second))
	nread, readAddr, err := listener.ReadFrom(buf)
	if err != nil {
		t.Errorf("ReadFrom() error = %v", err)
		return
	}

	if string(buf[:nread]) != "test message" {
		t.Errorf("ReadFrom() = %q, want %q", string(buf[:nread]), "test message")
	}

	if readAddr == nil {
		t.Error("ReadFrom() addr = nil, want non-nil")
	}
}

func TestUDPConn_ReadFromUDP(t *testing.T) {
	n := NewNet()

	// Create a UDP listener
	addr, err := n.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveUDPAddr() error = %v", err)
	}

	listener, err := n.ListenUDP("udp", addr)
	if err != nil {
		t.Fatalf("ListenUDP() error = %v", err)
	}
	defer listener.Close()

	listenAddr := listener.LocalAddr().(*UDPAddr)

	// Create a client conn
	client, err := n.DialUDP("udp", nil, listenAddr)
	if err != nil {
		t.Fatalf("DialUDP() error = %v", err)
	}
	defer client.Close()

	// Write from client
	message := []byte("udp test")
	client.Write(message)

	// Read on server using ReadFromUDP
	buf := make([]byte, 100)
	listener.SetReadDeadline(time.Now().Add(1 * time.Second))
	nread, addr, err := listener.ReadFromUDP(buf)
	if err != nil {
		t.Errorf("ReadFromUDP() error = %v", err)
		return
	}

	if string(buf[:nread]) != "udp test" {
		t.Errorf("ReadFromUDP() = %q, want %q", string(buf[:nread]), "udp test")
	}

	if addr == nil {
		t.Error("ReadFromUDP() addr = nil, want non-nil")
	}
}

func TestUDPConn_Deadlines(t *testing.T) {
	n := NewNet()

	// Create a UDP listener
	addr, err := n.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveUDPAddr() error = %v", err)
	}

	conn, err := n.ListenUDP("udp", addr)
	if err != nil {
		t.Fatalf("ListenUDP() error = %v", err)
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

func TestUDPConn_Buffers(t *testing.T) {
	n := NewNet()

	// Create a UDP listener
	addr, err := n.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveUDPAddr() error = %v", err)
	}

	conn, err := n.ListenUDP("udp", addr)
	if err != nil {
		t.Fatalf("ListenUDP() error = %v", err)
	}
	defer conn.Close()

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

func TestUDPConn_RemoteAddr(t *testing.T) {
	n := NewNet()

	// Create a UDP listener
	addr, err := n.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		t.Fatalf("ResolveUDPAddr() error = %v", err)
	}

	listener, err := n.ListenUDP("udp", addr)
	if err != nil {
		t.Fatalf("ListenUDP() error = %v", err)
	}
	defer listener.Close()

	listenAddr := listener.LocalAddr().(*UDPAddr)

	// Dial the listener (creates a connected UDP socket)
	conn, err := n.DialUDP("udp", nil, listenAddr)
	if err != nil {
		t.Fatalf("DialUDP() error = %v", err)
	}
	defer conn.Close()

	// Test RemoteAddr (should be set for connected UDP socket)
	remoteAddr := conn.RemoteAddr()
	if remoteAddr == nil {
		t.Error("RemoteAddr() = nil, want non-nil for connected UDP socket")
	}
}
