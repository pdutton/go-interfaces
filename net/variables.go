package net

import (
	"net"
)

var (
	IPv4bcast     = net.IPv4bcast
	IPv4allsys    = net.IPv4allsys
	IPv4allrouter = net.IPv4allrouter
	IPv4zero      = net.IPv4zero

	IPv6zero                   = net.IPv6zero
	IPv6unspecified            = net.IPv6unspecified
	IPv6loopback               = net.IPv6loopback
	IPv6interfacelocalallnodes = net.IPv6interfacelocalallnodes
	IPv6linklocalallnodes      = net.IPv6linklocalallnodes
	IPv6linklocalallrouters    = net.IPv6linklocalallrouters

	DefaultResolver = wrapResolver(net.DefaultResolver)

	ErrClosed           = net.ErrClosed
	ErrWriteToConnected = net.ErrWriteToConnected
)
