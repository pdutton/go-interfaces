package client

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type CheckRedirectFn = func(*http.Request, []*http.Request) error

// Client is an interface for the net/http.Client struct
type Client interface {
	CloseIdleConnections()
	Do(*http.Request) (Response, error)
	Get(string) (Response, error)
	Head(string) (Response, error)
	Post(string, string, io.Reader) (Response, error)
	PostForm(string, url.Values) (Response, error)
}

// ClientOption allows you to set options on a client in the NewClient constructor
type ClientOption func(cl *http.Client)

// Set the net/http.Client Transport
func WithTransport(t RoundTripper) ClientOption {
	return func(cl *http.Client) {
		cl.Transport = t
	}
}

// Set the net/http.Client CheckRedirect
func WithCheckRedirect(f func(*http.Request, []*http.Request) error) ClientOption {
	return func(cl *http.Client) {
		cl.CheckRedirect = f
	}
}

// Set the net/http.Client CookieJar
func WithCookieJar(j CookieJar) ClientOption {
	return func(cl *http.Client) {
		cl.Jar = j
	}
}

// WithJar is just an alias to WithCookieJar
var WithJar = WithCookieJar

// Set the net/http.Client Timeout
func WithTimeout(t time.Duration) ClientOption {
	return func(cl *http.Client) {
		cl.Timeout = t
	}
}

type clientFacade struct {
	realClient http.Client
}

// NewClient creates a Client with default values
func NewClient(options ...ClientOption) Client {
	var facade clientFacade

	for _, opt := range options {
		opt(&facade.realClient)
	}

	return facade
}

func (c clientFacade) CloseIdleConnections() {
	c.realClient.CloseIdleConnections()
}

func (c clientFacade) Do(req *http.Request) (Response, error) {
	resp, err := c.realClient.Do(req)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}

func (c clientFacade) Get(url string) (Response, error) {
	resp, err := c.realClient.Get(url)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}

func (c clientFacade) Head(url string) (Response, error) {
	resp, err := c.realClient.Head(url)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}

func (c clientFacade) Post(url string, contentType string, body io.Reader) (Response, error) {
	resp, err := c.realClient.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}

func (c clientFacade) PostForm(url string, data url.Values) (Response, error) {
	resp, err := c.realClient.PostForm(url, data)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}
