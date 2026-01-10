package net

import (
	"net"
	"os"
	"syscall"
	"time"
)

type UnixListener interface {
	Accept() (Conn, error)
	AcceptUnix() (UnixConn, error)
	Addr() Addr
	Close() error
	File() (*os.File, error)
	SetDeadline(t time.Time) error
	SetUnlinkOnClose(unlink bool)
	SyscallConn() (syscall.RawConn, error)
}

type unixListenerFacade struct {
	unixListener *net.UnixListener
}

func (_ netFacade) ListenUnix(network string, laddr *UnixAddr) (UnixListener, error) {
	ul, err := net.ListenUnix(network, laddr)
	return unixListenerFacade{unixListener: ul}, err
}

func (ul unixListenerFacade) Accept() (Conn, error) {
	return ul.unixListener.Accept()
}

func (ul unixListenerFacade) AcceptUnix() (UnixConn, error) {
	return ul.unixListener.AcceptUnix()
}

func (ul unixListenerFacade) Addr() Addr {
	return ul.unixListener.Addr()
}

func (ul unixListenerFacade) Close() error {
	return ul.unixListener.Close()
}

func (ul unixListenerFacade) File() (*os.File, error) {
	return ul.unixListener.File()
}

func (ul unixListenerFacade) SetDeadline(t time.Time) error {
	return ul.unixListener.SetDeadline(t)
}

func (ul unixListenerFacade) SetUnlinkOnClose(unlink bool) {
	ul.unixListener.SetUnlinkOnClose(unlink)
}

func (ul unixListenerFacade) SyscallConn() (syscall.RawConn, error) {
	return ul.unixListener.SyscallConn()
}
