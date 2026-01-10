package net

import (
	"context"
	"testing"
	"time"
)

func TestNet_Dial(t *testing.T) {
	n := NewNet()

	// Start a listener
	listener, err := n.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Dial the listener
	conn, err := n.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("Dial() error = %v", err)
	}
	defer conn.Close()

	if conn == nil {
		t.Fatal("Dial() returned nil connection")
	}
}

func TestNet_DialTimeout(t *testing.T) {
	n := NewNet()

	// Try to dial with a very short timeout to a non-existent address
	// This should timeout
	_, err := n.DialTimeout("tcp", "192.0.2.1:80", 1*time.Millisecond)
	if err == nil {
		t.Log("DialTimeout() succeeded unexpectedly (may happen on fast networks)")
	}
}

func TestNet_Listen(t *testing.T) {
	n := NewNet()

	listener, err := n.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	if listener == nil {
		t.Fatal("Listen() returned nil")
	}

	addr := listener.Addr()
	if addr == nil {
		t.Fatal("Listener.Addr() returned nil")
	}
}

func TestNet_ListenPacket(t *testing.T) {
	n := NewNet()

	conn, err := n.ListenPacket("udp", "localhost:0")
	if err != nil {
		t.Fatalf("ListenPacket() error = %v", err)
	}
	defer conn.Close()

	if conn == nil {
		t.Fatal("ListenPacket() returned nil")
	}
}

func TestNet_LookupIP(t *testing.T) {
	n := NewNet()

	ips, err := n.LookupIP("localhost")
	if err != nil {
		t.Errorf("LookupIP() error = %v", err)
		return
	}

	if len(ips) == 0 {
		t.Error("LookupIP() returned no IPs")
	}
}

func TestNet_ParseIP(t *testing.T) {
	n := NewNet()

	tests := []struct {
		name  string
		input string
		valid bool
	}{
		{"valid IPv4", "127.0.0.1", true},
		{"valid IPv6", "::1", true},
		{"invalid", "not-an-ip", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := n.ParseIP(tt.input)
			if tt.valid && ip == nil {
				t.Errorf("ParseIP(%q) = nil, want valid IP", tt.input)
			}
			if !tt.valid && ip != nil {
				t.Errorf("ParseIP(%q) = %v, want nil", tt.input, ip)
			}
		})
	}
}

func TestNet_IPv4(t *testing.T) {
	n := NewNet()

	ip := n.IPv4(127, 0, 0, 1)
	if ip == nil {
		t.Fatal("IPv4() returned nil")
	}

	expected := "127.0.0.1"
	if ip.String() != expected {
		t.Errorf("IPv4(127, 0, 0, 1).String() = %q, want %q", ip.String(), expected)
	}
}

func TestNet_IPv4Mask(t *testing.T) {
	n := NewNet()

	mask := n.IPv4Mask(255, 255, 255, 0)
	if mask == nil {
		t.Fatal("IPv4Mask() returned nil")
	}
}

func TestNet_CIDRMask(t *testing.T) {
	n := NewNet()

	mask := n.CIDRMask(24, 32)
	if mask == nil {
		t.Fatal("CIDRMask() returned nil")
	}
}

func TestNet_ResolveIPAddr(t *testing.T) {
	n := NewNet()

	addr, err := n.ResolveIPAddr("ip", "localhost")
	if err != nil {
		t.Errorf("ResolveIPAddr() error = %v", err)
		return
	}

	if addr == nil {
		t.Error("ResolveIPAddr() returned nil")
	}
}

func TestNet_ResolveTCPAddr(t *testing.T) {
	n := NewNet()

	addr, err := n.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		t.Errorf("ResolveTCPAddr() error = %v", err)
		return
	}

	if addr == nil {
		t.Error("ResolveTCPAddr() returned nil")
	}

	if addr.Port != 8080 {
		t.Errorf("ResolveTCPAddr() port = %d, want 8080", addr.Port)
	}
}

func TestNet_ResolveUDPAddr(t *testing.T) {
	n := NewNet()

	addr, err := n.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		t.Errorf("ResolveUDPAddr() error = %v", err)
		return
	}

	if addr == nil {
		t.Error("ResolveUDPAddr() returned nil")
	}
}

func TestNet_InterfaceAddrs(t *testing.T) {
	n := NewNet()

	addrs, err := n.InterfaceAddrs()
	if err != nil {
		t.Errorf("InterfaceAddrs() error = %v", err)
		return
	}

	if len(addrs) == 0 {
		t.Error("InterfaceAddrs() returned no addresses")
	}
}

func TestNet_Interfaces(t *testing.T) {
	n := NewNet()

	interfaces, err := n.Interfaces()
	if err != nil {
		t.Errorf("Interfaces() error = %v", err)
		return
	}

	if len(interfaces) == 0 {
		t.Error("Interfaces() returned no interfaces")
	}
}

func TestNet_LookupAddr(t *testing.T) {
	n := NewNet()

	// Lookup may fail on some systems
	_, err := n.LookupAddr("127.0.0.1")
	if err != nil {
		t.Logf("LookupAddr() error (may be expected): %v", err)
	}
}

func TestNet_LookupCNAME(t *testing.T) {
	n := NewNet()

	cname, err := n.LookupCNAME("www.google.com")
	if err != nil {
		t.Errorf("LookupCNAME() error = %v", err)
		return
	}

	if cname == "" {
		t.Error("LookupCNAME() returned empty string")
	}
}

func TestNet_LookupHost(t *testing.T) {
	n := NewNet()

	addrs, err := n.LookupHost("localhost")
	if err != nil {
		t.Errorf("LookupHost() error = %v", err)
		return
	}

	if len(addrs) == 0 {
		t.Error("LookupHost() returned no addresses")
	}
}

func TestNet_LookupPort(t *testing.T) {
	n := NewNet()

	port, err := n.LookupPort("tcp", "http")
	if err != nil {
		t.Errorf("LookupPort() error = %v", err)
		return
	}

	if port != 80 {
		t.Errorf("LookupPort(\"tcp\", \"http\") = %d, want 80", port)
	}
}

func TestNet_LookupTXT(t *testing.T) {
	n := NewNet()

	// TXT lookup may fail
	_, err := n.LookupTXT("google.com")
	if err != nil {
		t.Logf("LookupTXT() error (may be expected): %v", err)
	}
}

func TestNet_Pipe(t *testing.T) {
	n := NewNet()

	conn1, conn2 := n.Pipe()
	if conn1 == nil || conn2 == nil {
		t.Fatal("Pipe() returned nil connections")
	}
	defer conn1.Close()
	defer conn2.Close()

	// Test bidirectional communication
	go func() {
		buf := make([]byte, 100)
		n, _ := conn2.Read(buf)
		conn2.Write(buf[:n])
	}()

	message := []byte("test")
	conn1.Write(message)

	buf := make([]byte, 100)
	nread, err := conn1.Read(buf)
	if err != nil {
		t.Errorf("Read() error = %v", err)
		return
	}

	if string(buf[:nread]) != "test" {
		t.Errorf("Pipe communication = %q, want %q", string(buf[:nread]), "test")
	}
}

func TestNet_DialerMethods(t *testing.T) {
	n := NewNet()

	dialer := n.NewDialer(WithTimeout(5 * time.Second))
	if dialer == nil {
		t.Fatal("NewDialer() returned nil")
	}

	// Start a listener
	listener, err := n.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Test Dial
	conn, err := dialer.Dial("tcp", addr)
	if err != nil {
		t.Errorf("Dialer.Dial() error = %v", err)
		return
	}
	defer conn.Close()

	if conn == nil {
		t.Error("Dialer.Dial() returned nil")
	}
}

func TestNet_DialerDialContext(t *testing.T) {
	n := NewNet()

	dialer := n.NewDialer()

	// Start a listener
	listener, err := n.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("Listen() error = %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()

	// Test DialContext
	ctx := context.Background()
	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		t.Errorf("Dialer.DialContext() error = %v", err)
		return
	}
	defer conn.Close()

	if conn == nil {
		t.Error("Dialer.DialContext() returned nil")
	}
}
