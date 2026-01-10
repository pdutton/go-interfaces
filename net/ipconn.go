package net

import (
	"net"
	"os"
	"syscall"
	"time"
)

type IPConn interface {
	Close() error
	File() (f *os.File, err error)
	LocalAddr() Addr
	Read(b []byte) (int, error)
	ReadFrom(b []byte) (int, Addr, error)
	ReadFromIP(b []byte) (int, *IPAddr, error)
	ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)
	RemoteAddr() Addr
	SetDeadline(t time.Time) error
	SetReadBuffer(bytes int) error
	SetReadDeadline(t time.Time) error
	SetWriteBuffer(bytes int) error
	SetWriteDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
	Write(b []byte) (int, error)
	WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)
	WriteTo(b []byte, addr Addr) (int, error)
	WriteToIP(b []byte, addr *IPAddr) (int, error)
}

type ipConnFacade struct {
	ipConn *net.IPConn
}

func (_ netFacade) DialIP(network string, laddr, raddr *IPAddr) (IPConn, error) {
	ipc, err := net.DialIP(network, laddr, raddr)

	return ipConnFacade{ipConn: ipc}, err
}
func (_ netFacade) ListenIP(network string, laddr *IPAddr) (IPConn, error) {
	ipc, err := net.ListenIP(network, laddr)

	return ipConnFacade{ipConn: ipc}, err
}

func (ipc ipConnFacade) Close() error {
	return ipc.ipConn.Close()
}

func (ipc ipConnFacade) File() (f *os.File, err error) {
	return ipc.ipConn.File()
}

func (ipc ipConnFacade) LocalAddr() Addr {
	return ipc.ipConn.LocalAddr()
}

func (ipc ipConnFacade) Read(b []byte) (int, error) {
	return ipc.ipConn.Read(b)
}

func (ipc ipConnFacade) ReadFrom(b []byte) (int, Addr, error) {
	return ipc.ipConn.ReadFrom(b)
}

func (ipc ipConnFacade) ReadFromIP(b []byte) (int, *IPAddr, error) {
	return ipc.ipConn.ReadFromIP(b)
}

func (ipc ipConnFacade) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error) {
	return ipc.ipConn.ReadMsgIP(b, oob)
}

func (ipc ipConnFacade) RemoteAddr() Addr {
	return ipc.ipConn.RemoteAddr()
}

func (ipc ipConnFacade) SetDeadline(t time.Time) error {
	return ipc.ipConn.SetDeadline(t)
}

func (ipc ipConnFacade) SetReadBuffer(bytes int) error {
	return ipc.ipConn.SetReadBuffer(bytes)
}

func (ipc ipConnFacade) SetReadDeadline(t time.Time) error {
	return ipc.ipConn.SetReadDeadline(t)
}

func (ipc ipConnFacade) SetWriteBuffer(bytes int) error {
	return ipc.ipConn.SetWriteBuffer(bytes)
}

func (ipc ipConnFacade) SetWriteDeadline(t time.Time) error {
	return ipc.ipConn.SetWriteDeadline(t)
}

func (ipc ipConnFacade) SyscallConn() (syscall.RawConn, error) {
	return ipc.ipConn.SyscallConn()
}

func (ipc ipConnFacade) Write(b []byte) (int, error) {
	return ipc.ipConn.Write(b)
}

func (ipc ipConnFacade) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error) {
	return ipc.ipConn.WriteMsgIP(b, oob, addr)
}

func (ipc ipConnFacade) WriteTo(b []byte, addr Addr) (int, error) {
	return ipc.ipConn.WriteTo(b, addr)
}

func (ipc ipConnFacade) WriteToIP(b []byte, addr *IPAddr) (int, error) {
	return ipc.ipConn.WriteToIP(b, addr)
}
