package net

import (
	"net"
	"net/netip"
	"os"
	"syscall"
	"time"
)

type UDPConn interface {
	Close() error
	File() (*os.File, error)
	LocalAddr() Addr
	Read(b []byte) (int, error)
	ReadFrom(b []byte) (int, Addr, error)
	ReadFromUDP(b []byte) (int, *UDPAddr, error)
	ReadFromUDPAddrPort(b []byte) (int, netip.AddrPort, error)
	ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)
	ReadMsgUDPAddrPort(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error)
	RemoteAddr() Addr
	SetDeadline(t time.Time) error
	SetReadBuffer(bytes int) error
	SetReadDeadline(t time.Time) error
	SetWriteBuffer(bytes int) error
	SetWriteDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
	Write(b []byte) (int, error)
	WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)
	WriteMsgUDPAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error)
	WriteTo(b []byte, addr Addr) (int, error)
	WriteToUDP(b []byte, addr *UDPAddr) (int, error)
	WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error)
}

type udpConnFacade struct {
	udpConn *net.UDPConn
}

func (_ netFacade) DialUDP(network string, laddr, raddr *UDPAddr) (UDPConn, error) {
	c, err := net.DialUDP(network, laddr, raddr)
	return udpConnFacade{ udpConn: c }, err
}

func (_ netFacade) ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (UDPConn, error) {
	c, err := net.ListenMulticastUDP(network, ifi, gaddr)
	return udpConnFacade{ udpConn: c }, err
}

func (_ netFacade) ListenUDP(network string, laddr *UDPAddr) (UDPConn, error) {
	c, err := net.ListenUDP(network, laddr)
	return udpConnFacade{ udpConn: c }, err
}

func (f udpConnFacade) Close() error {
	return f.udpConn.Close()
}

func (f udpConnFacade) File() (*os.File, error) {
	return f.udpConn.File()
}

func (f udpConnFacade) LocalAddr() Addr {
	return f.udpConn.LocalAddr()
}

func (f udpConnFacade) Read(b []byte) (int, error) {
	return f.udpConn.Read(b)
}

func (f udpConnFacade) ReadFrom(b []byte) (int, Addr, error) {
	return f.udpConn.ReadFrom(b)
}

func (f udpConnFacade) ReadFromUDP(b []byte) (int, *UDPAddr, error) {
	return f.udpConn.ReadFromUDP(b)
}

func (f udpConnFacade) ReadFromUDPAddrPort(b []byte) (int, netip.AddrPort, error) {
	return f.udpConn.ReadFromUDPAddrPort(b)
}

func (f udpConnFacade) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error) {
	return f.udpConn.ReadMsgUDP(b, oob)
}

func (f udpConnFacade) ReadMsgUDPAddrPort(b, oob []byte) (n, oobn, flags int, addr netip.AddrPort, err error) {
	return f.udpConn.ReadMsgUDPAddrPort(b, oob)
}

func (f udpConnFacade) RemoteAddr() Addr {
	return f.udpConn.RemoteAddr()
}

func (f udpConnFacade) SetDeadline(t time.Time) error {
	return f.udpConn.SetDeadline(t)
}

func (f udpConnFacade) SetReadBuffer(bytes int) error {
	return f.udpConn.SetReadBuffer(bytes)
}

func (f udpConnFacade) SetReadDeadline(t time.Time) error {
	return f.udpConn.SetReadDeadline(t)
}

func (f udpConnFacade) SetWriteBuffer(bytes int) error {
	return f.udpConn.SetWriteBuffer(bytes)
}

func (f udpConnFacade) SetWriteDeadline(t time.Time) error {
	return f.udpConn.SetWriteDeadline(t)
}

func (f udpConnFacade) SyscallConn() (syscall.RawConn, error) {
	return f.udpConn.SyscallConn()
}

func (f udpConnFacade) Write(b []byte) (int, error) {
	return f.udpConn.Write(b)
}

func (f udpConnFacade) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error) {
	return f.udpConn.WriteMsgUDP(b, oob, addr)
}

func (f udpConnFacade) WriteMsgUDPAddrPort(b, oob []byte, addr netip.AddrPort) (n, oobn int, err error) {
	return f.udpConn.WriteMsgUDPAddrPort(b, oob, addr)
}

func (f udpConnFacade) WriteTo(b []byte, addr Addr) (int, error) {
	return f.udpConn.WriteTo(b, addr)
}

func (f udpConnFacade) WriteToUDP(b []byte, addr *UDPAddr) (int, error) {
	return f.udpConn.WriteToUDP(b, addr)
}

func (f udpConnFacade) WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error) {
	return f.udpConn.WriteToUDPAddrPort(b, addr)
}


