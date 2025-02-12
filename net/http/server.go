package http

import (
	"context"
	"net"
	"net/http"
)

// Server is an interface for the net/http.Server struct
type Server interface {
	Close() error
	ListenAndServe() error
	ListenAndServeTLS() error
	RegisterOnShutdown(func())
	Serve(net.Listener) error
	ServeTLS(net.Listener, string, string) error
	SetKeepAlivesEnabled(bool)
	Shutdown(context.Context) error
}

// ServerOption allows you to set options on a server in the NewServer constructor
type ServerOption func(svr *http.Server)

// Set the net/http.Server Addr
func WithAddr(addr string) ServerOption {
	return func(svr *http.Server) {
		svr.Addr = addr
	}
}

// Set the net/http.Server Handler
func WithHandler(handler Handler) ServerOption {
	return func(svr *http.Server) {
		svr.Handler = handler
	}
}

// Set the net/http.Server DisableGeneralOptionsHandler
func WithDisableGeneralOptionsHandler(v bool) ServerOption {
	return func(svr *http.Server) {
		svr.DisableGeneralOptionsHandler = v
	}
}

// Set the net/http.Server TLSConfig
func WithTLSConfig(cfg *tlsConfig) ServerOption {
	return func(svr *http.Server) {
		svr.TLSConfig = cfg
	}
}

// Set the net/http.Server ReadTimeout
func WithReadTimeout(t time.Duration) ServerOption {
	return func(svr *http.Server) {
		svr.ReadTimeout = t
	}
}

// Set the net/http.Server ReadHeaderTimeout
func WithReadHeaderTimeout(t time.Duration) ServerOption {
	return func(svr *http.Server) {
		svr.ReadHeaderTimeout = t
	}
}

// Set the net/http.Server WriteTimeout
func WithWriteTimeout(t time.Duration) ServerOption {
	return func(svr *http.Server) {
		svr.WriteTimeout = t
	}
}

// Set the net/http.Server IdleTimeout
func WithIdleTimeout(t time.Duration) ServerOption {
	return func(svr *http.Server) {
		svr.IdleTimeout = t
	}
}

// Set the net/http.Server MaxHeaderBytes
func WithMaxHeaderBytes(n int) ServerOption {
	return func(svr *http.Server) {
		svr.MaxHeaderBytes = n
	}
}

// Set the net/http.Server TLSNextProto
// TODO: Explore the possibility of replacing the function parameters with interfaces
func WithTLSNextProto(m map[string]func(*http.Server, *tls.Conn, http.Handler)) ServerOption {
	return func(svr *http.Server) {
		svr.TLSNextProto = m
	}
}

// Set the net/http.Server ConnState
// Note that the function will be called with the underlying connection, not an interface.
func WithConnState(f func(net.Conn, ConnState)) ServerOption {
	return func(svr *http.Server) {
		svr.ConnState = f
	}
}

// Set the net/http.Server ErrorLog
func WithErrorLog(l *log.Logger) ServerOption {
	return func(svr *http.Server) {
		svr.ErrorLog = l
	}
}

// Set the net/http.Server BaseContext
func WithBaseContext(f func(net.Listener) context.Context) ServerOption {
	return func(svr *http.Server) {
		svr.BaseContext = f
	}
}

// Set the net/http.Server ConnContext
func WithConnContext(f func(context.Context, net.Conn) context.Context) ServerOption {
	return func(svr *http.Server) {
		svr.ConnContext = f
	}
}

// Set the net/http.Server HTTP2
func WithHTTP2(cfg *HTTP2Config) ServerOption {
	return func(svr *http.Server) {
		svr.HTTP2 = cfg
	}
}

// Set the net/http.Server Protocols
func WithProtocols(p *Protocols) ServerOption {
	return func(svr *http.Server) {
		svr.Protocols = p
	}
}

type serverFacade struct {
	realServer http.Server
}

// NewServer creates a Server with default values
func NewServer(options ...ServerOption) Server {
	var facade serverFacade

	for _, opt := range options {
		opt(&facade.realServer)
	}

	return facade
}

func (s Server) Close() error {
	return s.realServer.Close()
}

func (s Server) ListenAndServe() error {
	return s.realServer.ListenAndServe()
}

func (s Server) ListenAndServeTLS() error {
	return s.realServer.ListenAndServeTLS()
}

func (s Server) RegisterOnShutdown(f func()) {
	return s.realServer.RegisterOnShutdown(f)
}

func (s Server) Serve(l net.Listener) error {
	return s.realServer.Serve(l)
}

func (s Server) ServeTLS(l net.Listener, certFile, keyFile string) error {
	return s.realServer.ServeTLS(l, certFile, keyFile)
}

func (s Server) SetKeepAlivesEnabled(v bool) {
	return s.realServer.SetKeepAlivesEnabled(v)
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.realServer.Shutdown(ctx)
}


