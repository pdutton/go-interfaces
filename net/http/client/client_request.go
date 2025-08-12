package http

import (
	"bufio"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

// Request provides an interface to a net/http.Request for use on
// a client making an http request.
type Request interface {
	Write(io.Writer) error
	WriteProxy(io.Writer) error

	RealRequest() *http.Request

	// Is Clone useful?
	// Clone(context.Context) *Request

	// Set these in the constructor instead:
	// SetBasicAuth(username, password string)
	// SetPathValue(name, value string)
	// WithContext(context.Context) Request
}

// RequestOption allows you to set options on a request in the NewRequest constructor
type RequestOption func(req *http.Request)

// Set the net/http.Request Header map
func WithHeaders(headers Header) RequestOption {
	return func(req *http.Request) {
		req.Header = headers
	}
}

// Set a an individual net/http.Request Header
func WithHeader(name string, values ...string) RequestOption {
	return func(req *http.Request) {
		// Initialize the Header map if it is nil
		if req.Header == nil {
			req.Header = make(map[string][]string)
		}
		req.Header[name] = values
	}
}

// Set the net/http.Request Body
func WithBody(r io.ReadCloser) RequestOption {
	return func(req *http.Request) {
		req.Body = r
	}
}

// Set the net/http.Request GetBody
func WithGetBody(f func() (io.ReadCloser, error)) RequestOption {
	return func(req *http.Request) {
		req.GetBody = f
	}
}

// Set the net/http.Request ContentLength
func WithContentLength(l int64) RequestOption {
	return func(req *http.Request) {
		req.ContentLength = l
	}
}

// Set the net/http.Request TransferEncoding
func WithTransferEncoding(encodings []string) RequestOption {
	return func(req *http.Request) {
		req.TransferEncoding = encodings
	}
}

// Set the net/http.Request Close
func WithClose(v bool) RequestOption {
	return func(req *http.Request) {
		req.Close = v
	}
}

// Set the net/http.Request Host
func WithHost(host string) RequestOption {
	return func(req *http.Request) {
		req.Host = host
	}
}

// Set the net/http.Request Form
func WithForm(form url.Values) RequestOption {
	return func(req *http.Request) {
		req.Form = form
	}
}

// Set the net/http.Request PostForm
func WithPostForm(form url.Values) RequestOption {
	return func(req *http.Request) {
		req.PostForm = form
	}
}

// Set the net/http.Request MultipartForm
func WithMultipartForm(form *multipart.Form) RequestOption {
	return func(req *http.Request) {
		req.MultipartForm = form
	}
}

// Set the net/http.RequestForm Trailer
func WithTrailer(t Header) RequestOption {
	return func(req *http.Request) {
		req.Trailer = t
	}
}

// Set the net/http.RequestForm RemoteAddr
func WithRemoteAddr(addr string) RequestOption {
	return func(req *http.Request) {
		req.RemoteAddr = addr
	}
}

// Set the net/http.RequestForm RequestURI
func WithRequestURI(uri string) RequestOption {
	return func(req *http.Request) {
		req.RequestURI = uri
	}
}

type requestFacade struct {
	realRequest *http.Request
}

func newRequest(req *http.Request) requestFacade {
	return requestFacade{
		realRequest: req,
	}
}

func (_ httpFacade) NewRequest(method, url string, body io.Reader, options ...RequestOption) (Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	var facade = requestFacade{
		realRequest: req,
	}

	facade.setOptions(options...)

	return facade, nil
}

func (_ httpFacade) NewRequestWithContext(ctx context.Context, method, url string, body io.Reader, options ...RequestOption) (Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	var facade = requestFacade{
		realRequest: req,
	}

	facade.setOptions(options...)

	return facade, nil
}

func (f requestFacade) RealRequest() *http.Request {
	return f.realRequest
}

func (f requestFacade) setOptions(options ...RequestOption) {
	for _, opt := range options {
		opt(f.realRequest)
	}
}

// ReadRequest is a "constructor" that constructs a Request from a Reader
func ReadRequest(r *bufio.Reader) (Request, error) {
	var facade requestFacade

	req, err := http.ReadRequest(r)
	if err != nil {
		return facade, err
	}

	facade.realRequest = req

	return facade, nil
}

func (r requestFacade) Write(w io.Writer) error {
	return r.realRequest.Write(w)
}

func (r requestFacade) WriteProxy(w io.Writer) error {
	return r.realRequest.WriteProxy(w)
}
