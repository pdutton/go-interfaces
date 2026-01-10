package net

import (
	"context"
	"net"
	"syscall"
	"time"
)

type ListenConfig interface {
	Listen(ctx context.Context, network, address string) (Listener, error)
	ListenPacket(ctx context.Context, network, address string) (PacketConn, error)
	MultipathTCP() bool
	// StMultipathTCP(bool)
}

type ListenConfigOption func(*net.ListenConfig)

func WithControlLC(f func(string, string, syscall.RawConn) error) ListenConfigOption {
	return func(lc *net.ListenConfig) {
		lc.Control = f
	}
}

func WithKeepAliveLC(d time.Duration) ListenConfigOption {
	return func(lc *net.ListenConfig) {
		lc.KeepAlive = d
	}
}

func WithKeepAliveConfigLC(kac KeepAliveConfig) ListenConfigOption {
	return func(lc *net.ListenConfig) {
		lc.KeepAliveConfig = kac
	}
}

type listenConfigFacade struct {
	listenConfig *net.ListenConfig
}

func (_ netFacade) NewListenConfig(options ...ListenConfigOption) ListenConfig {
	var lc net.ListenConfig

	for _, opt := range options {
		opt(&lc)
	}

	return listenConfigFacade{
		listenConfig: &lc,
	}
}

func (lc listenConfigFacade) Listen(ctx context.Context, network, address string) (Listener, error) {
	return lc.listenConfig.Listen(ctx, network, address)
}

func (lc listenConfigFacade) ListenPacket(ctx context.Context, network, address string) (PacketConn, error) {
	return lc.listenConfig.ListenPacket(ctx, network, address)
}

func (lc listenConfigFacade) MultipathTCP() bool {
	return lc.listenConfig.MultipathTCP()
}
