package http

import (
	"crypto/tls"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

// TODO: Determine if mocking the server request is really useful.  If so, finish:

// Request provides an interface to a net/http.Request for use on
// a server receiving an http request.
type Request interface {
	Method() string
	URL() *url.URL
	Proto() string
	ProtoMajor() int
	ProtoMinor() int
	Header() Header
	Body() io.ReadCloser
	GetBody() func() (io.ReadCloser, error)
	ContentLength() int64
	TransferEncoding() []string
	Host() string
	Form() url.Values
	PostForm() url.Values
	MultipartForm() *multipart.Form
	Trailer() Header
	RemoteAddr() string
	RequestURI() string
	TLS() *tls.ConnectionState
	Pattern() string
}

type serverRequestFacade struct {
	realRequest *http.Request
}
