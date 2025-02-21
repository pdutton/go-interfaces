package http

import (
	"bufio"
	"io"
	"net/http"
	"net/url"
)

// HTTP is an interface for the functions in the net/http package
type HTTP interface {
	// Response constructors:
	Get(string) (Response, error)
	Head(string) (Response, error)
	Post(string, string, io.Reader) (Response, error)
	PostForm(string, url.Values) (Response, error)
	ReadResponse(*bufio.Reader, *http.Request) (Response, error)

	CanonicalHeaderKey(string) string
	StatusText(int) string
}

type httpFacade struct{}

// NewHTTP creates a new HTTP instance.
func NewHTTP() HTTP {
	return httpFacade{}
}

func (_ httpFacade) Get(url string) (Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}

func (_ httpFacade) Head(url string) (Response, error) {
	resp, err := http.Head(url)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}

func (_ httpFacade) Post(url string, contentType string, body io.Reader) (Response, error) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}

func (_ httpFacade) PostForm(url string, values url.Values) (Response, error) {
	resp, err := http.PostForm(url, values)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}

func (_ httpFacade) ReadResponse(r *bufio.Reader, req *http.Request) (Response, error) {
	resp, err := http.ReadResponse(r, req)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}

func (_ httpFacade) CanonicalHeaderKey(s string) string {
	return http.CanonicalHeaderKey(s)
}

func (_ httpFacade) StatusText(code int) string {
	return http.StatusText(code)
}
