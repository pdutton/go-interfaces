package net

import (
	"context"
	"testing"
	"time"
)

func TestNewListenConfig(t *testing.T) {
	n := NewNet()
	lc := n.NewListenConfig()
	if lc == nil {
		t.Fatal("NewListenConfig() returned nil")
	}
}

func TestListenConfig_WithOptions(t *testing.T) {
	n := NewNet()
	lc := n.NewListenConfig(
		WithKeepAliveLC(30 * time.Second),
	)

	if lc == nil {
		t.Fatal("NewListenConfig() with options returned nil")
	}
}

func TestListenConfig_Listen(t *testing.T) {
	n := NewNet()
	lc := n.NewListenConfig()

	// Listen on an ephemeral port
	listener, err := lc.Listen(context.Background(), "tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	if listener == nil {
		t.Fatal("Listen() returned nil listener")
	}

	addr := listener.Addr()
	if addr == nil {
		t.Fatal("Listener.Addr() returned nil")
	}
}

func TestListenConfig_ListenPacket(t *testing.T) {
	n := NewNet()
	lc := n.NewListenConfig()

	// Listen for UDP packets
	conn, err := lc.ListenPacket(context.Background(), "udp", "localhost:0")
	if err != nil {
		t.Fatalf("ListenPacket() error = %v", err)
	}
	defer conn.Close()

	if conn == nil {
		t.Fatal("ListenPacket() returned nil connection")
	}

	addr := conn.LocalAddr()
	if addr == nil {
		t.Fatal("PacketConn.LocalAddr() returned nil")
	}
}

func TestListenConfig_WithContext(t *testing.T) {
	n := NewNet()
	lc := n.NewListenConfig()

	// Create a context that will be canceled
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	listener, err := lc.Listen(ctx, "tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() with context error = %v", err)
	}
	defer listener.Close()

	if listener == nil {
		t.Fatal("Listen() with context returned nil")
	}
}
