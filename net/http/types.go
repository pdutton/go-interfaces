package http

import (
	"net/http"
)

// This file contains aliases for simple types in net/http that don't need
// an interface either because they are already an interface, or they are 
// unlikely to require mocking. They are specified in this package so that
// calling code doesn't need to import the net/http package directly.

type ConnState      = http.ConnState
type Cookie         = http.Cookie
type CookieJar      = http.CookieJar
type Dir            = http.Dir
type File           = http.File
type FileSystem     = http.FileSystem
type Flusher        = http.Flusher
// HTTP2Config is being added in 1.24.0.  When it's a little more baked, I'll add it here:
// type HTTP2Config    = http.HTTP2Config
type Handler        = http.Handler
type HandlerFunc    = http.HandlerFunc
type Header         = http.Header
type Hijacker       = http.Hijacker
type MaxBytesError  = http.MaxBytesError
// Protocols is being added in 1.24.0.  When it's a little more baked, I'll add it here:
// type Protocols      = http.Protocols
type PushOptions    = http.PushOptions
type Pusher         = http.Pusher
type ResponseWriter = http.ResponseWriter
type RoundTripper   = http.RoundTripper
type SameSite       = http.SameSite






