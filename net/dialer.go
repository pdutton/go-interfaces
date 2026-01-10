package net

import (
	"context"
	"net"
	"syscall"
	"time"
)

type Dialer interface {
	Dial(network, address string) (Conn, error)
	DialContext(ctx context.Context, network, address string) (Conn, error)
	MultipathTCP() bool
	// SetMultipathTCP(use)   // Use WithSetMultipathTCP
}

type dialerFacade struct {
	dialer *net.Dialer
}

type DialerOption func(*net.Dialer)

func WithTimeout(d time.Duration) DialerOption {
	return func(dia *net.Dialer) {
		dia.Timeout = d
	}
}

func WithDeadline(t time.Time) DialerOption {
	return func(dia *net.Dialer) {
		dia.Deadline = t
	}
}

func WithLocalAddr(addr Addr) DialerOption {
	return func(dia *net.Dialer) {
		dia.LocalAddr = addr
	}
}

func WithDualStack() DialerOption {
	return func(dia *net.Dialer) {
		dia.DualStack = true
	}
}

func WithFallbackDelay(d time.Duration) DialerOption {
	return func(dia *net.Dialer) {
		dia.FallbackDelay = d
	}
}

func WithKeepAlive(d time.Duration) DialerOption {
	return func(dia *net.Dialer) {
		dia.KeepAlive = d
	}
}

func WithKeepAliveConfig(cfg KeepAliveConfig) DialerOption {
	return func(dia *net.Dialer) {
		dia.KeepAliveConfig = cfg
	}
}

func WithCancel(c <-chan struct{}) DialerOption {
	return func(dia *net.Dialer) {
		dia.Cancel = c
	}
}

func WithResolver(r Resolver) DialerOption {
	return func(dia *net.Dialer) {
		dia.Resolver = r.GetUnderlyingResolver()
	}
}

func WithControl(f func(string, string, syscall.RawConn) error) DialerOption {
	return func(dia *net.Dialer) {
		dia.Control = f
	}
}

func WithControlContext(f func(context.Context, string, string, syscall.RawConn) error) DialerOption {
	return func(dia *net.Dialer) {
		dia.ControlContext = f
	}
}

func WithSetMultipathTCP(b bool) DialerOption {
	return func(dia *net.Dialer) {
		dia.SetMultipathTCP(b)
	}
}

func (_ netFacade) NewDialer(options ...DialerOption) Dialer {
	var dialer net.Dialer

	for _, opt := range options {
		opt(&dialer)
	}

	return dialerFacade{dialer: &dialer}
}

func (d dialerFacade) Dial(network, address string) (Conn, error) {
	return d.dialer.Dial(network, address)
}

func (d dialerFacade) DialContext(ctx context.Context, network, address string) (Conn, error) {
	return d.dialer.DialContext(ctx, network, address)
}

func (d dialerFacade) MultipathTCP() bool {
	return d.dialer.MultipathTCP()
}
