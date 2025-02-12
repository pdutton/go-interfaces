package http

import (
	"bufio"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

// ClientRequest provides an interface to a net/http.Request for use on
// a client making an http request.
type ClientRequest interface {
	Write(io.Writer) error
	WriteProxy(io.Writer) error

	// Is Clone useful?
	// Clone(context.Context) *Request

	// Set these in the constructor instead:
	// SetBasicAuth(username, password string)
	// SetPathValue(name, value string)
	// WithContext(context.Context) Request
}

// ClientRequestOption allows you to set options on a request in the NewClientRequest constructor
type ClientRequestOption func(req *http.Request)

// Set the net/http.Request Header map
func WithHeaders(headers Header) ClientRequestOption {
	return func(req *http.Request) {
		req.Header = headers
	}
}

// Set a an individual net/http.Request Header
func WithHeader(name string, values ...string) ClientRequestOption {
	return func(req *http.Request) {
		// Initialize the Header map if it is nil
		if req.Header == nil {
			req.Header = make(map[string][]string)
		}
		req.Header[name] = values
	}
}

// Set the net/http.Request Body
func WithBody(r io.ReadCloser) ClientRequestOption {
	return func(req *http.Request) {
		req.Body = r
	}
}

// Set the net/http.Request GetBody
func WithGetBody(f func() (io.ReadCloser, error)) ClientRequestOption {
	return func(req *http.Request) {
		req.GetBody = f
	}
}

// Set the net/http.Request ContentLength
func WithContentLength(l int64) ClientRequestOption {
	return func(req *http.Request) {
		req.ContentLength = l
	}
}

// Set the net/http.Request TransferEncoding
func WithTransferEncoding(encodings []string) ClientRequestOption {
	return func(req *http.Request) {
		req.TransferEncoding = encodings
	}
}

// Set the net/http.Request Close
func WithClose(v bool) ClientRequestOption {
	return func(req *http.Request) {
		req.Close = v
	}
}

// Set the net/http.Request Host
func WithHost(host string) ClientRequestOption {
	return func(req *http.Request) {
		req.Host = host
	}
}

// Set the net/http.Request Form
func WithForm(form url.Values) ClientRequestOption {
	return func(req *http.Request) {
		req.Form = form
	}
}

// Set the net/http.Request PostForm
func WithPostForm(form url.Values) ClientRequestOption {
	return func(req *http.Request) {
		req.PostForm = form
	}
}

// Set the net/http.Request MultipartForm
func WithMultipartForm(form *multipart.Form) ClientRequestOption {
	return func(req *http.Request) {
		req.MultipartForm = form
	}
}

// Set the net/http.RequestForm Trailer
func WithTrailer(t Header) ClientRequestOption {
	return func(req *http.Request) {
		req.Trailer = t
	}
}

// Set the net/http.RequestForm RemoteAddr
func WithRemoteAddr(addr string) ClientRequestOption {
	return func(req *http.Request) {
		req.RemoteAddr = addr
	}
}

// Set the net/http.RequestForm RequestURI
func WithRequestURI(uri string) ClientRequestOption {
	return func(req *http.Request) {
		req.RequestURI = uri
	}
}

type clientRequestFacade struct {
	realRequest *http.Request
}

func newClientRequest(req *http.Request) clientRequestFacade {
	return clientRequestFacade{
		realRequest: req,
	}
}

func NewClientRequest(method, url string, body io.Reader, options ...ClientRequestOption) (ClientRequest, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	var facade = clientRequestFacade{
		realRequest: req,
	}

	facade.setOptions(options...)

	return facade, nil
}

func NewClientRequestWithContext(ctx context.Context, method, url string, body io.Reader, options ...ClientRequestOption) (ClientRequest, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	var facade = clientRequestFacade{
		realRequest: req,
	}

	facade.setOptions(options...)

	return facade, nil
}

func (f clientRequestFacade) setOptions(options ...ClientRequestOption) {
	for _, opt := range options {
		opt(f.realRequest)
	}
}

// ReadRequest is a "constructor" that constructs a Request from a Reader
func ReadRequest(r *bufio.Reader) (ClientRequest, error) {
	var facade clientRequestFacade

	req, err := http.ReadRequest(r)
	if err != nil {
		return facade, err
	}

	facade.realRequest = req

	return facade, nil
}

func (r clientRequestFacade) Write(w io.Writer) error {
	return r.realRequest.Write(w)
}

func (r clientRequestFacade) WriteProxy(w io.Writer) error {
	return r.realRequest.WriteProxy(w)
}





