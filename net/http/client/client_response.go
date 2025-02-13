package client

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
)

// Response is an interface to a Response received by the client.
type Response interface {
	// Simple pass-through methods:
	Cookies() []*Cookie
	Location() (*url.URL, error)
	ProtoAtLeast(int, int) bool
	Write(io.Writer) error

	// Methods allowing access to member variables:
	Status() string
	StatusCode() int
	Proto() string
	ProtoMajor() int
	ProtoMinor() int
	Header() Header
	Body() io.ReadCloser
	ContentLength() int64
	TransferEncoding() []string
	Close() bool
	Uncompressed() bool
	Trailer() Header
	Request() Request
	TLS() *tls.ConnectionState
}

type responseFacade struct {
	realResponse *http.Response
}

func newResponse(resp *http.Response) responseFacade {
	return responseFacade{
		realResponse: resp,
	}
}

func (r responseFacade) Cookies() []*Cookie {
	return r.realResponse.Cookies()
}

func (r responseFacade) Location() (*url.URL, error) {
	return r.realResponse.Location()
}

func (r responseFacade) ProtoAtLeast(major, minor int) bool {
	return r.realResponse.ProtoAtLeast(major, minor)
}

func (r responseFacade) Write(w io.Writer) error {
	return r.realResponse.Write(w)
}

func (r responseFacade) Status() string {
	return r.realResponse.Status
}

func (r responseFacade) StatusCode() int {
	return r.realResponse.StatusCode
}

func (r responseFacade) Proto() string {
	return r.realResponse.Proto
}

func (r responseFacade) ProtoMajor() int {
	return r.realResponse.ProtoMajor
}

func (r responseFacade) ProtoMinor() int {
	return r.realResponse.ProtoMinor
}

func (r responseFacade) Header() Header {
	return r.realResponse.Header
}

func (r responseFacade) Body() io.ReadCloser {
	return r.realResponse.Body
}

func (r responseFacade) ContentLength() int64 {
	return r.realResponse.ContentLength
}

func (r responseFacade) TransferEncoding() []string {
	return r.realResponse.TransferEncoding
}

func (r responseFacade) Close() bool {
	return r.realResponse.Close
}

func (r responseFacade) Uncompressed() bool {
	return r.realResponse.Uncompressed
}

func (r responseFacade) Trailer() Header {
	return r.realResponse.Trailer
}

func (r responseFacade) Request() Request {
	return newRequest(r.realResponse.Request)
}

func (r responseFacade) TLS() *tls.ConnectionState {
	return r.realResponse.TLS
}
