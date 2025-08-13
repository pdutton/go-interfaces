package http

import (
	"net/http"
)

// This file contains aliases for simple types in net/http that don't need
// an interface either because they are already an interface, or they are
// unlikely to require mocking. They are specified in this package so that
// calling code doesn't need to import the net/http package directly.

type ConnState = http.ConnState
type Cookie = http.Cookie
type CookieJar = http.CookieJar

type HTTP2Config = http.HTTP2Config
type Header = http.Header

type Protocols = http.Protocols
type RoundTripper = http.RoundTripper
type SameSite = http.SameSite
type Transport = http.Transport
