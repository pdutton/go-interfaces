package net

import (
	"context"
	"net"
	"net/netip"
)

type Resolver interface {
	LookupAddr(ctx context.Context, addr string) ([]string, error)
	LookupCNAME(ctx context.Context, host string) (string, error)
	LookupHost(ctx context.Context, host string) (addrs []string, err error)
	LookupIP(ctx context.Context, network, host string) ([]IP, error)
	LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error)
	LookupMX(ctx context.Context, name string) ([]*MX, error)
	LookupNS(ctx context.Context, name string) ([]*NS, error)
	LookupNetIP(ctx context.Context, network, host string) ([]netip.Addr, error)
	LookupPort(ctx context.Context, network, service string) (port int, err error)
	LookupSRV(ctx context.Context, service, port, name string) (string, []*SRV, error)
	LookupTXT(ctx context.Context, name string) ([]string, error)
}

type ResolverOption func(res *net.Resolver)

func WithPreferGo() ResolverOption {
	return func(res *net.Resolver) {
		res.PreferGo = true
	}
}

func WithStrictErrors() ResolverOption {
	return func(res *net.Resolver) {
		res.StrictErrors = true
	}
}

func WithDialFunc(fn func(context.Context, string, string) (Conn, error)) ResolverOption {
	return func(res *net.Resolver) {
		res.Dial = fn
	}
}

type resolverFacade struct {
	resolver *net.Resolver
}

func (_ netFacade) NewResolver(options ...ResolverOption) Resolver {
	var realResolver net.Resolver

	for _, opt := range options {
		opt(&realResolver)
	}

	return resolverFacade{ resolver: &realResolver }
}

func wrapResolver(res *net.Resolver) Resolver {
	return resolverFacade{ resolver: res }
}

func (r resolverFacade) LookupAddr(ctx context.Context, addr string) ([]string, error) {
	return r.resolver.LookupAddr(ctx, addr)
}

func (r resolverFacade) LookupCNAME(ctx context.Context, host string) (string, error) {
	return r.resolver.LookupCNAME(ctx, host)
}

func (r resolverFacade) LookupHost(ctx context.Context, host string) (addrs []string, err error) {
	return r.resolver.LookupHost(ctx, host)
}

func (r resolverFacade) LookupIP(ctx context.Context, network, host string) ([]IP, error) {
	return r.resolver.LookupIP(ctx, network, host)
}

func (r resolverFacade) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error) {
	return r.resolver.LookupIPAddr(ctx, host)
}

func (r resolverFacade) LookupMX(ctx context.Context, name string) ([]*MX, error) {
	return r.resolver.LookupMX(ctx, name)
}

func (r resolverFacade) LookupNS(ctx context.Context, name string) ([]*NS, error) {
	return r.resolver.LookupNS(ctx, name)
}

func (r resolverFacade) LookupNetIP(ctx context.Context, network, host string) ([]netip.Addr, error) {
	return r.resolver.LookupNetIP(ctx, network, host)
}

func (r resolverFacade) LookupPort(ctx context.Context, network, service string) (port int, err error) {
	return r.resolver.LookupPort(ctx, network, service)
}

func (r resolverFacade) LookupSRV(ctx context.Context, service, port, name string) (string, []*SRV, error) {
	return r.resolver.LookupSRV(ctx, service, port, name)
}

func (r resolverFacade) LookupTXT(ctx context.Context, name string) ([]string, error) {
	return r.resolver.LookupTXT(ctx, name)
}

