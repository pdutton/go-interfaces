package net

import (
	"context"
	"net"
	"net/netip"
	"os"
	"time"
)

type Net interface {
	JoinHostPort(host, port string) string
	LookupAddr(addr string) (names []string, err error)
	LookupCNAME(host string) (cname string, err error)
	LookupHost(host string) (addrs []string, err error)
	LookupPort(network, service string) (port int, err error)
	LookupTXT(name string) ([]string, error)
	ParseCIDR(s string) (IP, *IPNet, error)
	Pipe() (Conn, Conn)
	SplitHostPort(hostport string) (host, port string, err error)

	// Constructors:
	InterfaceAddrs() ([]Addr, error)
	Dial(network, address string) (Conn, error)
	DialTimeout(network, address string, timeout time.Duration) (Conn, error)
	FileConn(f *os.File) (Conn, error)
	IPv4(a, b, c, d byte) IP
	LookupIP(host string) ([]IP, error)
	ParseIP(s string) IP
	ResolveIPAddr(network, address string) (*IPAddr, error)
	DialIP(network string, laddr, raddr *IPAddr) (IPConn, error)
	ListenIP(network string, laddr *IPAddr) (IPConn, error)
	CIDRMask(ones, bits int) IPMask
	IPv4Mask(a, b, c, d byte) IPMask
	InterfaceByIndex(index int) (*Interface, error)
	InterfaceByName(name string) (*Interface, error)
	Interfaces() ([]Interface, error)
	FileListener(f *os.File) (Listener, error)
	Listen(network, address string) (Listener, error)
	LookupMX(name string) ([]*MX, error)
	LookupNS(name string) ([]*NS, error)
	FilePacketConn(f *os.File) (PacketConn, error)
	ListenPacket(network, address string) (PacketConn, error)
	LookupSRV(service, proto, name string) (string, []*SRV, error)
	ResolveTCPAddr(network, address string) (*TCPAddr, error)
	TCPAddrFromAddrPort(addr netip.AddrPort) *TCPAddr
	DialTCP(network string, laddr, raddr *TCPAddr) (TCPConn, error)
	ListenTCP(network string, laddr *TCPAddr) (TCPListener, error)
	ResolveUDPAddr(network, address string) (*UDPAddr, error)
	DialUDP(network string, laddr, raddr *UDPAddr) (UDPConn, error)
	ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (UDPConn, error)
	ListenUDP(network string, laddr *UDPAddr) (UDPConn, error)
	ResolveUnixAddr(network, address string) (*UnixAddr, error)
	DialUnix(network string, laddr, raddr *UnixAddr) (UnixConn, error)
	ListenUnixgram(network string, laddr *UnixAddr) (UnixConn, error)
	ListenUnix(network string, laddr *UnixAddr) (UnixListener, error)
}

type netFacade struct {
}

func NewNet() netFacade {
	return netFacade{}
}

func (_ netFacade) JoinHostPort(host, port string) string {
	return net.JoinHostPort(host, port)
}

func (_ netFacade) LookupAddr(addr string) (names []string, err error) {
	return DefaultResolver.LookupAddr(context.Background(), addr)
}

func (_ netFacade) LookupCNAME(host string) (cname string, err error) {
	return DefaultResolver.LookupCNAME(context.Background(), host)
}

func (_ netFacade) LookupHost(host string) (addrs []string, err error) {
	return DefaultResolver.LookupHost(context.Background(), host)
}

func (_ netFacade) LookupPort(network, service string) (port int, err error) {
	return DefaultResolver.LookupPort(context.Background(), network, service)
}

func (_ netFacade) LookupTXT(name string) ([]string, error) {
	return DefaultResolver.LookupTXT(context.Background(), name)
}

func (_ netFacade) ParseCIDR(s string) (IP, *IPNet, error) {
	return net.ParseCIDR(s)
}

func (_ netFacade) Pipe() (Conn, Conn) {
	return net.Pipe()
}

func (_ netFacade) SplitHostPort(hostport string) (host, port string, err error) {
	return net.SplitHostPort(hostport)
}

func (_ netFacade) InterfaceAddrs() ([]Addr, error) {
	return net.InterfaceAddrs()
}

func (_ netFacade) Dial(network, address string) (Conn, error) {
	return net.Dial(network, address)
}

func (_ netFacade) DialTimeout(network, address string, timeout time.Duration) (Conn, error) {
	return net.DialTimeout(network, address, timeout)
}

func (_ netFacade) FileConn(f *os.File) (Conn, error) {
	return net.FileConn(f)
}

func (_ netFacade) IPv4(a, b, c, d byte) IP {
	return net.IPv4(a, b, c, d)
}

func (_ netFacade) LookupIP(host string) ([]IP, error) {
	return net.LookupIP(host)
}

func (_ netFacade) ParseIP(s string) IP {
	return net.ParseIP(s)
}

func (_ netFacade) ResolveIPAddr(network, address string) (*IPAddr, error) {
	return net.ResolveIPAddr(network, address)
}

func (_ netFacade) CIDRMask(ones, bits int) IPMask {
	return net.CIDRMask(ones, bits)
}

func (_ netFacade) IPv4Mask(a, b, c, d byte) IPMask {
	return net.IPv4Mask(a, b, c, d)
}

func (_ netFacade) InterfaceByIndex(index int) (*Interface, error) {
	return net.InterfaceByIndex(index)
}

func (_ netFacade) InterfaceByName(name string) (*Interface, error) {
	return net.InterfaceByName(name)
}

func (_ netFacade) Interfaces() ([]Interface, error) {
	return net.Interfaces()
}

func (_ netFacade) FileListener(f *os.File) (Listener, error) {
	return net.FileListener(f)
}

func (_ netFacade) Listen(network, address string) (Listener, error) {
	return net.Listen(network, address)
}

func (_ netFacade) LookupMX(name string) ([]*MX, error) {
	return net.LookupMX(name)
}

func (_ netFacade) LookupNS(name string) ([]*NS, error) {
	return net.LookupNS(name)
}

func (_ netFacade) FilePacketConn(f *os.File) (PacketConn, error) {
	return net.FilePacketConn(f)
}

func (_ netFacade) ListenPacket(network, address string) (PacketConn, error) {
	return net.ListenPacket(network, address)
}

func (_ netFacade) LookupSRV(service, proto, name string) (string, []*SRV, error) {
	return net.LookupSRV(service, proto, name)
}

func (_ netFacade) ResolveTCPAddr(network, address string) (*TCPAddr, error) {
	return net.ResolveTCPAddr(network, address)
}

func (_ netFacade) TCPAddrFromAddrPort(addr netip.AddrPort) *TCPAddr {
	return net.TCPAddrFromAddrPort(addr)
}

func (_ netFacade) ResolveUDPAddr(network, address string) (*UDPAddr, error) {
	return net.ResolveUDPAddr(network, address)
}

func (_ netFacade) ResolveUnixAddr(network, address string) (*UnixAddr, error) {
	return net.ResolveUnixAddr(network, address)
}
