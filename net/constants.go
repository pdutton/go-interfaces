package net

import (
	"net"
)

const (
	IPv4len = net.IPv4len
	IPv6len = net.IPv6len

	FlagUp           = net.FlagUp
	FlagBroadcast    = net.FlagBroadcast
	FlagLoopback     = net.FlagLoopback
	FlagPointToPoint = net.FlagPointToPoint
	FlagMulticast    = net.FlagMulticast
	FlagRunning      = net.FlagRunning
)
