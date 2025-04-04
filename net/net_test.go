package net

import (
    "reflect"
    "testing"
)

func TestNewNet(t *testing.T) {
    var impl = NewNet()

    var intf Net = impl

    t.Logf(`intf: %v`, intf)
}

func TestNet_JoinHostPort(t *testing.T) {
    var impl = NewNet()

    if hp := impl.JoinHostPort(`example.com`, `80`); hp != `example.com:80` {
        t.Errorf(`unexpected host-port: %s`, hp)
    }
}

func TestNet_ParseCIDR_IPv4(t *testing.T) {
    var impl = NewNet()

    var cidr = `192.168.0.12/24`
    var expectIP4 = IP{ 192, 168, 0, 12 }
    var expectIP16 = IP{
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 192, 168, 0, 12 }
    var expectMask = IPNet{
        IP{ 192, 168, 0, 0 },
        IPMask{ 0xff, 0xff, 0xff, 0x00 },
    }

    ip, ipnet, err := impl.ParseCIDR(cidr)
    if err != nil {
        t.Errorf(`unexpected error: %v`, err)
    }
    if !reflect.DeepEqual(ip, expectIP4) && !reflect.DeepEqual(ip, expectIP16) {
        t.Errorf(`unexpected ip: %v`, []byte(ip))
    }
    if !reflect.DeepEqual(ipnet, &expectMask) {
        t.Errorf(`unexpected mask: %v % x`, []byte(ipnet.IP), []byte(ipnet.Mask))
    }
}

func TestNet_ParseCIDR_IPv6(t *testing.T) {
    var impl = NewNet()

    var cidr = `2001:db8:a0b:12f0::1/32`

    var expectIP = IP{ 
        0x20, 0x01, 0x0d, 0xb8, 0x0a, 0x0b, 0x12, 0xf0,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
    }
    var expectMask = IPNet{
        IP{ 0x20, 0x01, 0x0d, 0xb8, 0x00, 0x00, 0x00, 0x00,
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
        IPMask{ 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00,
            0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
    }

    ip, ipnet, err := impl.ParseCIDR(cidr)
    if err != nil {
        t.Errorf(`unexpected error: %v`, err)
    }
    if !reflect.DeepEqual(ip, expectIP) {
        t.Errorf(`unexpected ip: % x`, []byte(ip))
    }
    if !reflect.DeepEqual(ipnet, &expectMask) {
        t.Errorf(`unexpected mask: % x % x`, []byte(ipnet.IP), []byte(ipnet.Mask))
    }
}

func TestNet_SplitHost(t *testing.T) {
    var impl = NewNet()

    host, port, err := impl.SplitHostPort(`example.com:443`)
    if err != nil {
        t.Errorf(`unexpected error: %v`, err)
    }
    if host != `example.com` {
        t.Errorf(`unexpected host: %s`, host)
    }
    if port != `443` {
        t.Errorf(`unexpected port: %s`, port)
    }
}

