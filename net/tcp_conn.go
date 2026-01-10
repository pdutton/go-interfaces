package net

import (
	"io"
	"net"
	"os"
	"syscall"
	"time"
)

type TCPConn interface {
	Close() error
	CloseRead() error
	CloseWrite() error
	File() (*os.File, error)
	LocalAddr() Addr
	MultipathTCP() (bool, error)
	Read(b []byte) (int, error)
	ReadFrom(r io.Reader) (int64, error)
	RemoteAddr() Addr
	SetDeadline(t time.Time) error
	SetKeepAlive(keepalive bool) error
	SetKeepAliveConfig(config KeepAliveConfig) error
	SetKeepAlivePeriod(d time.Duration) error
	SetLinger(sec int) error
	SetNoDelay(noDelay bool) error
	SetReadBuffer(bytes int) error
	SetReadDeadline(t time.Time) error
	SetWriteBuffer(bytes int) error
	SetWriteDeadline(t time.Time) error
	SyscallConn() (syscall.RawConn, error)
	Write(b []byte) (int, error)
	WriteTo(w io.Writer) (int64, error)
}

type tcpConnFacade struct {
	tcpConn *net.TCPConn
}

func (_ netFacade) DialTCP(network string, laddr, raddr *TCPAddr) (TCPConn, error) {
	tcpc, err := net.DialTCP(network, laddr, raddr)
	return tcpConnFacade{tcpConn: tcpc}, err
}

func (tcpc tcpConnFacade) Close() error {
	return tcpc.tcpConn.Close()
}

func (tcpc tcpConnFacade) CloseRead() error {
	return tcpc.tcpConn.CloseRead()
}

func (tcpc tcpConnFacade) CloseWrite() error {
	return tcpc.tcpConn.CloseWrite()
}

func (tcpc tcpConnFacade) File() (*os.File, error) {
	return tcpc.tcpConn.File()
}

func (tcpc tcpConnFacade) LocalAddr() Addr {
	return tcpc.tcpConn.LocalAddr()
}

func (tcpc tcpConnFacade) MultipathTCP() (bool, error) {
	return tcpc.tcpConn.MultipathTCP()
}

func (tcpc tcpConnFacade) Read(b []byte) (int, error) {
	return tcpc.tcpConn.Read(b)
}

func (tcpc tcpConnFacade) ReadFrom(r io.Reader) (int64, error) {
	return tcpc.tcpConn.ReadFrom(r)
}

func (tcpc tcpConnFacade) RemoteAddr() Addr {
	return tcpc.tcpConn.RemoteAddr()
}

func (tcpc tcpConnFacade) SetDeadline(t time.Time) error {
	return tcpc.tcpConn.SetDeadline(t)
}

func (tcpc tcpConnFacade) SetKeepAlive(keepalive bool) error {
	return tcpc.tcpConn.SetKeepAlive(keepalive)
}

func (tcpc tcpConnFacade) SetKeepAliveConfig(config KeepAliveConfig) error {
	return tcpc.tcpConn.SetKeepAliveConfig(config)
}

func (tcpc tcpConnFacade) SetKeepAlivePeriod(d time.Duration) error {
	return tcpc.tcpConn.SetKeepAlivePeriod(d)
}

func (tcpc tcpConnFacade) SetLinger(sec int) error {
	return tcpc.tcpConn.SetLinger(sec)
}

func (tcpc tcpConnFacade) SetNoDelay(noDelay bool) error {
	return tcpc.tcpConn.SetNoDelay(noDelay)
}

func (tcpc tcpConnFacade) SetReadBuffer(bytes int) error {
	return tcpc.tcpConn.SetReadBuffer(bytes)
}

func (tcpc tcpConnFacade) SetReadDeadline(t time.Time) error {
	return tcpc.tcpConn.SetReadDeadline(t)
}

func (tcpc tcpConnFacade) SetWriteBuffer(bytes int) error {
	return tcpc.tcpConn.SetWriteBuffer(bytes)
}

func (tcpc tcpConnFacade) SetWriteDeadline(t time.Time) error {
	return tcpc.tcpConn.SetWriteDeadline(t)
}

func (tcpc tcpConnFacade) SyscallConn() (syscall.RawConn, error) {
	return tcpc.tcpConn.SyscallConn()
}

func (tcpc tcpConnFacade) Write(b []byte) (int, error) {
	return tcpc.tcpConn.Write(b)
}

func (tcpc tcpConnFacade) WriteTo(w io.Writer) (int64, error) {
	return tcpc.tcpConn.WriteTo(w)
}
