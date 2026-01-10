package net

import (
	"context"
	"testing"
)

func TestResolver_LookupHost(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	addrs, err := r.LookupHost(context.Background(), "localhost")
	if err != nil {
		t.Errorf("LookupHost() error = %v", err)
		return
	}

	if len(addrs) == 0 {
		t.Error("LookupHost() returned no addresses")
	}
}

func TestResolver_LookupIP(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	ips, err := r.LookupIP(context.Background(), "ip", "localhost")
	if err != nil {
		t.Errorf("LookupIP() error = %v", err)
		return
	}

	if len(ips) == 0 {
		t.Error("LookupIP() returned no IPs")
	}
}

func TestResolver_LookupIPAddr(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	addrs, err := r.LookupIPAddr(context.Background(), "localhost")
	if err != nil {
		t.Errorf("LookupIPAddr() error = %v", err)
		return
	}

	if len(addrs) == 0 {
		t.Error("LookupIPAddr() returned no addresses")
	}
}

func TestResolver_LookupNetIP(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	ips, err := r.LookupNetIP(context.Background(), "ip", "localhost")
	if err != nil {
		t.Errorf("LookupNetIP() error = %v", err)
		return
	}

	if len(ips) == 0 {
		t.Error("LookupNetIP() returned no IPs")
	}
}

func TestResolver_LookupAddr(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	// Use a known localhost IP
	names, err := r.LookupAddr(context.Background(), "127.0.0.1")
	if err != nil {
		// LookupAddr may fail on some systems, so just check it doesn't panic
		t.Logf("LookupAddr() error (may be expected): %v", err)
		return
	}

	if len(names) > 0 {
		t.Logf("LookupAddr() returned: %v", names)
	}
}

func TestResolver_LookupCNAME(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	// Use a well-known domain
	cname, err := r.LookupCNAME(context.Background(), "www.google.com")
	if err != nil {
		t.Errorf("LookupCNAME() error = %v", err)
		return
	}

	if cname == "" {
		t.Error("LookupCNAME() returned empty string")
	}
}

func TestResolver_LookupPort(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	port, err := r.LookupPort(context.Background(), "tcp", "http")
	if err != nil {
		t.Errorf("LookupPort() error = %v", err)
		return
	}

	if port != 80 {
		t.Errorf("LookupPort() = %d, want 80", port)
	}
}

func TestResolver_LookupTXT(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	// TXT lookup may fail on some networks
	_, err := r.LookupTXT(context.Background(), "google.com")
	if err != nil {
		t.Logf("LookupTXT() error (may be expected): %v", err)
	}
}

func TestResolver_LookupMX(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	// MX lookup may fail on some networks
	_, err := r.LookupMX(context.Background(), "google.com")
	if err != nil {
		t.Logf("LookupMX() error (may be expected): %v", err)
	}
}

func TestResolver_LookupNS(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	// NS lookup may fail on some networks
	_, err := r.LookupNS(context.Background(), "google.com")
	if err != nil {
		t.Logf("LookupNS() error (may be expected): %v", err)
	}
}

func TestResolver_LookupSRV(t *testing.T) {
	n := NewNet()
	r := n.NewResolver()

	// SRV lookup may fail on some networks
	_, _, err := r.LookupSRV(context.Background(), "xmpp-server", "tcp", "google.com")
	if err != nil {
		t.Logf("LookupSRV() error (may be expected): %v", err)
	}
}

func TestResolver_WithOptions(t *testing.T) {
	// Test creating resolver with options
	n := NewNet()
	r := n.NewResolver(WithPreferGo())

	if r == nil {
		t.Fatal("NewResolver() with options returned nil")
	}

	// Test it still works
	addrs, err := r.LookupHost(context.Background(), "localhost")
	if err != nil {
		t.Errorf("LookupHost() with PreferGo error = %v", err)
		return
	}

	if len(addrs) == 0 {
		t.Error("LookupHost() with PreferGo returned no addresses")
	}
}
