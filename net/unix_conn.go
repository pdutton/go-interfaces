package net

import (
	"net"
	"os"
	"syscall"
	"time"
)

type UnixConn interface {
	Close() error
	CloseRead() error
	CloseWrite() error
	File() (*os.File, error)
	LocalAddr() Addr
	Read(b []byte) (int, error)
	ReadFrom(b []byte) (int, Addr, error)
	ReadFromUnix(b []byte) (int, *UnixAddr, error)
	ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)
	RemoteAddr() Addr
	SetDeadline(t time.Time) error
	SetReadBuffer(bytes int) error
	SetReadDeadline(t time.Time) error
	SetWriteBuffer(bytes int) error
	SetWriteDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
	Write(b []byte) (int, error)
	WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)
	WriteTo(b []byte, addr Addr) (int, error)
	WriteToUnix(b []byte, addr *UnixAddr) (int, error)
}

type unixConnFacade struct {
	unixConn *net.UnixConn
}

func (_ netFacade) DialUnix(network string, laddr, raddr *UnixAddr) (UnixConn, error) {
	uc, err := net.DialUnix(network, laddr, raddr)
	return unixConnFacade{ unixConn: uc }, err
}

func (_ netFacade) ListenUnixgram(network string, laddr *UnixAddr) (UnixConn, error) {
	uc, err := net.ListenUnixgram(network, laddr)
	return unixConnFacade{ unixConn: uc }, err
}

func (uc unixConnFacade) Close() error {
	return uc.unixConn.Close()
}

func (uc unixConnFacade) CloseRead() error {
	return uc.unixConn.CloseRead()
}

func (uc unixConnFacade) CloseWrite() error {
	return uc.unixConn.CloseWrite()
}

func (uc unixConnFacade) File() (*os.File, error) {
	return uc.unixConn.File()
}

func (uc unixConnFacade) LocalAddr() Addr {
	return uc.unixConn.LocalAddr()
}

func (uc unixConnFacade) Read(b []byte) (int, error) {
	return uc.unixConn.Read(b)
}

func (uc unixConnFacade) ReadFrom(b []byte) (int, Addr, error) {
	return uc.unixConn.ReadFrom(b)
}

func (uc unixConnFacade) ReadFromUnix(b []byte) (int, *UnixAddr, error) {
	return uc.unixConn.ReadFromUnix(b)
}

func (uc unixConnFacade) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error) {
	return uc.unixConn.ReadMsgUnix(b, oob)
}

func (uc unixConnFacade) RemoteAddr() Addr {
	return uc.unixConn.RemoteAddr()
}

func (uc unixConnFacade) SetDeadline(t time.Time) error {
	return uc.unixConn.SetDeadline(t)
}

func (uc unixConnFacade) SetReadBuffer(bytes int) error {
	return uc.unixConn.SetReadBuffer(bytes)
}

func (uc unixConnFacade) SetReadDeadline(t time.Time) error {
	return uc.unixConn.SetReadDeadline(t)
}

func (uc unixConnFacade) SetWriteBuffer(bytes int) error {
	return uc.unixConn.SetWriteBuffer(bytes)
}

func (uc unixConnFacade) SetWriteDeadline(t time.Time) error {
	return uc.unixConn.SetWriteDeadline(t)
}

func (uc unixConnFacade) SyscallConn() (syscall.RawConn, error) {
	return uc.unixConn.SyscallConn()
}

func (uc unixConnFacade) Write(b []byte) (int, error) {
	return uc.unixConn.Write(b)
}

func (uc unixConnFacade) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error) {
	return uc.unixConn.WriteMsgUnix(b, oob, addr)
}

func (uc unixConnFacade) WriteTo(b []byte, addr Addr) (int, error) {
	return uc.unixConn.WriteTo(b, addr)
}

func (uc unixConnFacade) WriteToUnix(b []byte, addr *UnixAddr) (int, error) {
	return uc.unixConn.WriteToUnix(b, addr)
}


