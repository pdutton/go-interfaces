package net

import (
	"net"
	"os"
	"syscall"
	"time"
)

type TCPListener interface {
	Accept() (Conn, error)
	AcceptTCP() (TCPConn, error)
	Addr() Addr
	Close() error
	File() (*os.File, error)
	SetDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
}

type tcpListenerFacade struct {
	tcpListener *net.TCPListener
}

func (_ netFacade) ListenTCP(network string, laddr *TCPAddr) (TCPListener, error) {
	tcpl, err := net.ListenTCP(network, laddr)

	return tcpListenerFacade{ tcpListener: tcpl }, err
}

func (f tcpListenerFacade) Accept() (Conn, error) {
	return f.tcpListener.Accept()
}

func (f tcpListenerFacade) AcceptTCP() (TCPConn, error) {
	return f.tcpListener.AcceptTCP()
}

func (f tcpListenerFacade) Addr() Addr {
	return f.tcpListener.Addr()
}

func (f tcpListenerFacade) Close() error {
	return f.tcpListener.Close()
}

func (f tcpListenerFacade) File() (*os.File, error) {
	return f.tcpListener.File()
}

func (f tcpListenerFacade) SetDeadline(t time.Time) error {
	return f.tcpListener.SetDeadline(t)
}

func (f tcpListenerFacade) SyscallConn() (syscall.RawConn, error) {
	return f.tcpListener.SyscallConn()
}

