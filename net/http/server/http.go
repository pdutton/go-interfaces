package server

import (
	"io"
	"io/fs"
	"net"
	"net/http"
	"net/url"
	"time"
)

// HTTP is an interface for the functions in the net/http package
// when used in a server
type HTTP interface {
	CanonicalHeaderKey(string) string
	DetectContentType([]byte) string
	Error(ResponseWriter, string, int)
	Handle(string, http.Handler)
	HandleFunc(string, func(ResponseWriter, *http.Request))
	ListenAndServe(string, http.Handler) error
	ListenAndServeTLS(string, string, string, http.Handler) error
	MaxBytesReader(ResponseWriter, io.ReadCloser, int64) io.ReadCloser
	NotFound(ResponseWriter, *http.Request)
	ParseHTTPVersion(string) (int, int, bool)
	ParseTime(string) (time.Time, error)
	ProxyFromEnvironment(*http.Request) (*url.URL, error)
	ProxyURL(*url.URL) func(*http.Request) (*url.URL, error)
	Redirect(ResponseWriter, *http.Request, string, int)
	Serve(net.Listener, http.Handler) error
	ServeContent(ResponseWriter, *http.Request, string, time.Time, io.ReadSeeker)
	ServeFile(ResponseWriter, *http.Request, string)
	ServeFileFS(ResponseWriter, *http.Request, fs.FS, string)
	ServeTLS(net.Listener, http.Handler, string, string) error
	SetCookie(ResponseWriter, *Cookie)
	StatusText(int) string
}

type httpFacade struct{}

// NewHTTP creates a new HTTP instance.
func NewHTTP() HTTP {
	return httpFacade{}
}

func (_ httpFacade) CanonicalHeaderKey(s string) string {
	return http.CanonicalHeaderKey(s)
}

func (_ httpFacade) DetectContentType(data []byte) string {
	return http.DetectContentType(data)
}

func (_ httpFacade) Error(w ResponseWriter, error string, code int) {
	http.Error(w, error, code)
}

func (_ httpFacade) Handle(pattern string, handler http.Handler) {
	http.Handle(pattern, handler)
}

func (_ httpFacade) HandleFunc(pattern string, handler func(ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, handler)
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

func (_ httpFacade) NotFound(w ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func (_ httpFacade) ParseHTTPVersion(vers string) (int, int, bool) {
	return http.ParseHTTPVersion(vers)
}

func (_ httpFacade) ParseTime(text string) (time.Time, error) {
	return http.ParseTime(text)
}

func (_ httpFacade) ProxyFromEnvironment(req *http.Request) (*url.URL, error) {
	return http.ProxyFromEnvironment(req)
}

func (_ httpFacade) ProxyURL(fixedURL *url.URL) func(*http.Request) (*url.URL, error) {
	return http.ProxyURL(fixedURL)
}

func (_ httpFacade) Redirect(w ResponseWriter, r *http.Request, url string, code int) {
	http.Redirect(w, r, url, code)
}

func (_ httpFacade) Serve(l net.Listener, h http.Handler) error {
	return http.Serve(l, h)
}

func (_ httpFacade) ServeContent(w ResponseWriter, req *http.Request, name string, modtime time.Time, content io.ReadSeeker) {
	http.ServeContent(w, req, name, modtime, content)
}

func (_ httpFacade) ServeFile(w ResponseWriter, r *http.Request, name string) {
	http.ServeFile(w, r, name)
}

func (_ httpFacade) ServeFileFS(w ResponseWriter, r *http.Request, fsys fs.FS, name string) {
	http.ServeFileFS(w, r, fsys, name)
}

func (_ httpFacade) ServeTLS(l net.Listener, handler http.Handler, certFile string, keyFile string) error {
	return http.ServeTLS(l, handler, certFile, keyFile)
}

func (_ httpFacade) SetCookie(w ResponseWriter, cookie *Cookie) {
	http.SetCookie(w, cookie)
}

func (_ httpFacade) StatusText(code int) string {
	return http.StatusText(code)
}
