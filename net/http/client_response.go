package http

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
)

// ClientResponse is an interface to a Response received by the client.
type ClientResponse interface {
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
	Request() ClientRequest
	TLS() *tls.ConnectionState
}

type clientResponseFacade struct {
	realResponse *http.Response
}

func newClientResponse(resp *http.Response) clientResponseFacade {
	return clientResponseFacade{
		realResponse: resp,
	}
}

func (r clientResponseFacade) Cookies() []*Cookie {
	return r.realResponse.Cookies()
}

func (r clientResponseFacade) Location() (*url.URL, error) {
	return r.realResponse.Location()
}

func (r clientResponseFacade) ProtoAtLeast(major, minor int) bool {
	return r.realResponse.ProtoAtLeast(major, minor)
}

func (r clientResponseFacade) Write(w io.Writer) error {
	return r.realResponse.Write(w)
}


func (r clientResponseFacade) Status() string {
	return r.realResponse.Status
}

func (r clientResponseFacade) StatusCode() int {
	return r.realResponse.StatusCode
}

func (r clientResponseFacade) Proto() string {
	return r.realResponse.Proto
}

func (r clientResponseFacade) ProtoMajor() int {
	return r.realResponse.ProtoMajor
}

func (r clientResponseFacade) ProtoMinor() int {
	return r.realResponse.ProtoMinor
}

func (r clientResponseFacade) Header() Header {
	return r.realResponse.Header
}

func (r clientResponseFacade) Body() io.ReadCloser {
	return r.realResponse.Body
}

func (r clientResponseFacade) ContentLength() int64 {
	return r.realResponse.ContentLength
}

func (r clientResponseFacade) TransferEncoding() []string {
	return r.realResponse.TransferEncoding
}

func (r clientResponseFacade) Close() bool {
	return r.realResponse.Close
}

func (r clientResponseFacade) Uncompressed() bool {
	return r.realResponse.Uncompressed
}

func (r clientResponseFacade) Trailer() Header {
	return r.realResponse.Trailer
}

func (r clientResponseFacade) Request() ClientRequest {
	return newClientRequest(r.realResponse.Request)
}

func (r clientResponseFacade) TLS() *tls.ConnectionState {
	return r.realResponse.TLS
}



