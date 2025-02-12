package http

import (
	"net"
	"net/http"
	"net/url"
	"io"
	"io/fs"
	"time"
)

// HTTP is an interface for the functions in the net/http package 
type HTTP interface {
	// ClientResponse constructors:
	Get(string) (ClientResponse, error)
	Head(string) (ClientResponse, error)
	Post(string, string, io.Reader) (ClientResponse, error)
	PostForm(string, url.Values) (ClientResponse, error)
	ReadResponse(*bufio.Reader, ClientRequest) (ClientResponse, error)


	CanonicalHeaderKey(string) string
	DetectContentType([]byte) string
	Error(ResponseWriter, string, int)
	Handle(string, http.Handler)
	HandlerFunc(string, func(ResponseWriter, http.Request))
	ListenAndServe(string, http.Handler) error
	ListenAndServeTLS(string, string, string, http.Handler) error
	MaxBytesReader(ResponseWriter, io.ReadCloser, int64) io.ReadCloser
	NotFound(ResponseWriter, *Request)
	ParseHTTPVersion(string) (int, int, bool)
	ParseTime(string) (time.Time, error)
	ProxyFromEnvironment(*Request) (*url.URL, error)
	ProxyURL(*url.URL) func(*Request) (*url.URL, error)
	Redirect(ResponseWriter, *Request, string, int)
	Serve(net.Listener, http.Handler) error
	ServeContent(ResponseWriter, *Request, string, time.Time, io.ReadSeeker)
	ServeFile(ResponseWriter, *Request, string)
	ServeFileFS(ResponseWriter, *Request, fs.FS, string)
	ServeTLS(net.Listener, http.Handler, string, string) error
	SetCookie(ResponseWriter, *Cookie)
	StatusText(int) string
}

type httpFacade struct {}

// NewHTTP creates a new HTTP instance.
func NewHTTP() HTTP {
	return httpFacade{}
}

func (_ httpFacade) Get(url string) (ClientResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return newClientResponse(resp), nil
}

func (_ httpFacade) Head(url string) (ClientResponse, error) {
	resp, err := http.Head(url)
	if err != nil {
		return nil, err
	}

	return newClientResponse(resp), nil
}

func (_ httpFacade) Post(url string, contentType string, body io.Reader) (ClientResponse, error) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}

	return newClientResponse(resp), nil
}

func (_ httpFacade) PostForm(url string, values url.Values) (ClientResponse, error) {
	resp, err := http.PostForm(url, values)
	if err != nil {
		return nil, err
	}

	return newClientResponse(resp), nil
}

func (_ httpFacade) ReadResponse(r *bufio.Reader, req ClientRequest) (ClientResponse, error) {
	resp, err := http.ReadResponse(r, req)
	if err != nil {
		return nil, err
	}

	return newClientResonse(resp), nil
}

func (_ httpFacade) CanonicalHeaderKey(s string) string {
	return http.CanonicalHeaderKey(s)
}

func (_ httpFacade) DetectContentType(data []byte) string {
	return http.DetectContentType(data)
}

func (_ httpFacade) Error(w ResponseWriter, error string, code int) {
	return http.Error(w, error, code)
}

func (_ httpFacade) Handle(pattern string, handler http.Handler) {
	return http.Handle(pattern, handler)
}

func (_ httpFacade) HandlerFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	return http.HandleFunc(pattern, handler)
}

func (_ httpFacade) ListenAndServe(addr string, handler http.Handler) error {
	return http.ListenAndServe(addr, handler)
}

func (_ httpFacade) ListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, handler)
}

func (_ httpFacade) MaxBytesReader(w ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser {
	return http.MaxBytesReader(w, r, n)
}

func (_ httpFacade) NotFound(w ResponseWriter, r *Request) {
	return http.NotFound(w, r)
}

func (_ httpFacade) ParseHTTPVersion(vers string) (int, int, bool) {
	return http.ParseHTTPVersion(vers)
}

func (_ httpFacade) ParseTime(text string) (time.Time, error) {
	return http.ParseTime(text)
}

func (_ httpFacade) ProxyFromEnvironment(req *Request) (*url.URL, error) {
	return http.ProxyFromEnvironment(req)
}

func (_ httpFacade) ProxyURL(fixedURL *url.URL) func(*Request) (*url.URL, error) {
	return http.ProxyURL(fixedURL)
}

func (_ httpFacade) Redirect(w ResponseWriter, r *Request, url string, code int) {
	return http.Redirect(w, r, url, code)
}

func (_ httpFacade) Serve(l net.Listener, h http.Handler) error {
	return http.Serve(l, h)
}

func (_ httpFacade) ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker) {
	return http.ServeContent(w, req, name, modtime, content)
}

func (_ httpFacade) ServeFile(w ResponseWriter, r *Request, name string) {
	return http.ServeFile(w, r, name)
}

func (_ httpFacade) ServeFileFS(w ResponseWriter, r *Request, fsys fs.FS, name string) {
	return http.ServeFileFS(w, r, fsys, name)
}

func (_ httpFacade) ServeTLS(l net.Listener, handler http.Handler, certFile string, keyFile string) error {
	return http.ServeTLS(l, handler, certFile, keyFile)
}

func (_ httpFacade) SetCookie(w ResponseWriter, cookie *Cookie) {
	return http.SetCookie(w, cookie)
}

func (_ httpFacade) StatusText(code int) string {
	return http.StatusText(code)
}


