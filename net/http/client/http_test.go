package http

import (
	"bytes"
	"io"
	stdhttp "net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestNewHTTP(t *testing.T) {
	h := NewHTTP()
	_ = h
}

func TestHTTP_Get(t *testing.T) {
	// Create test server
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		if r.Method != "GET" {
			t.Errorf("Method = %s, want GET", r.Method)
		}
		w.Write([]byte("test response"))
	}))
	defer server.Close()

	h := NewHTTP()
	resp, err := h.Get(server.URL)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	defer resp.Body().Close()

	body, err := io.ReadAll(resp.Body())
	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}

	if string(body) != "test response" {
		t.Errorf("Response body = %q, want %q", string(body), "test response")
	}

	if resp.StatusCode() != 200 {
		t.Errorf("StatusCode = %d, want 200", resp.StatusCode())
	}
}

func TestHTTP_Head(t *testing.T) {
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		if r.Method != "HEAD" {
			t.Errorf("Method = %s, want HEAD", r.Method)
		}
		w.Header().Set("X-Test-Header", "test-value")
	}))
	defer server.Close()

	h := NewHTTP()
	resp, err := h.Head(server.URL)
	if err != nil {
		t.Fatalf("Head() error = %v", err)
	}
	defer resp.Body().Close()

	if resp.StatusCode() != 200 {
		t.Errorf("StatusCode = %d, want 200", resp.StatusCode())
	}

	if resp.Header().Get("X-Test-Header") != "test-value" {
		t.Errorf("Header = %q, want %q", resp.Header().Get("X-Test-Header"), "test-value")
	}
}

func TestHTTP_Post(t *testing.T) {
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		if r.Method != "POST" {
			t.Errorf("Method = %s, want POST", r.Method)
		}

		body, _ := io.ReadAll(r.Body)
		if string(body) != "test data" {
			t.Errorf("Body = %q, want %q", string(body), "test data")
		}

		w.Write([]byte("posted"))
	}))
	defer server.Close()

	h := NewHTTP()
	resp, err := h.Post(server.URL, "text/plain", bytes.NewReader([]byte("test data")))
	if err != nil {
		t.Fatalf("Post() error = %v", err)
	}
	defer resp.Body().Close()

	body, _ := io.ReadAll(resp.Body())
	if string(body) != "posted" {
		t.Errorf("Response body = %q, want %q", string(body), "posted")
	}
}

func TestHTTP_PostForm(t *testing.T) {
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		if r.Method != "POST" {
			t.Errorf("Method = %s, want POST", r.Method)
		}

		r.ParseForm()
		if r.FormValue("key") != "value" {
			t.Errorf("Form value = %q, want %q", r.FormValue("key"), "value")
		}

		w.Write([]byte("form posted"))
	}))
	defer server.Close()

	h := NewHTTP()
	formData := url.Values{
		"key": []string{"value"},
	}

	resp, err := h.PostForm(server.URL, formData)
	if err != nil {
		t.Fatalf("PostForm() error = %v", err)
	}
	defer resp.Body().Close()

	body, _ := io.ReadAll(resp.Body())
	if string(body) != "form posted" {
		t.Errorf("Response body = %q, want %q", string(body), "form posted")
	}
}

func TestHTTP_NewRequest(t *testing.T) {
	h := NewHTTP()

	req, err := h.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatalf("NewRequest() error = %v", err)
	}

	if req == nil {
		t.Fatal("NewRequest() returned nil")
	}

	realReq := req.RealRequest()
	if realReq.Method != "GET" {
		t.Errorf("Method = %s, want GET", realReq.Method)
	}

	if realReq.URL.String() != "http://example.com" {
		t.Errorf("URL = %s, want http://example.com", realReq.URL.String())
	}
}

func TestHTTP_NewClient(t *testing.T) {
	h := NewHTTP()

	client := h.NewClient()
	if client == nil {
		t.Fatal("NewClient() returned nil")
	}
}

func TestClient_Do(t *testing.T) {
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		w.Write([]byte("client do"))
	}))
	defer server.Close()

	h := NewHTTP()
	client := h.NewClient()

	req, _ := h.NewRequest("GET", server.URL, nil)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}
	defer resp.Body().Close()

	body, _ := io.ReadAll(resp.Body())
	if string(body) != "client do" {
		t.Errorf("Response body = %q, want %q", string(body), "client do")
	}
}

func TestClient_Get(t *testing.T) {
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		w.Write([]byte("client get"))
	}))
	defer server.Close()

	h := NewHTTP()
	client := h.NewClient()

	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	defer resp.Body().Close()

	body, _ := io.ReadAll(resp.Body())
	if string(body) != "client get" {
		t.Errorf("Response body = %q, want %q", string(body), "client get")
	}
}

func TestClient_Head(t *testing.T) {
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		if r.Method != "HEAD" {
			t.Errorf("Method = %s, want HEAD", r.Method)
		}
	}))
	defer server.Close()

	h := NewHTTP()
	client := h.NewClient()

	resp, err := client.Head(server.URL)
	if err != nil {
		t.Fatalf("Head() error = %v", err)
	}
	defer resp.Body().Close()

	if resp.StatusCode() != 200 {
		t.Errorf("StatusCode = %d, want 200", resp.StatusCode())
	}
}

func TestClient_Post(t *testing.T) {
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		body, _ := io.ReadAll(r.Body)
		w.Write(body)
	}))
	defer server.Close()

	h := NewHTTP()
	client := h.NewClient()

	resp, err := client.Post(server.URL, "text/plain", strings.NewReader("post data"))
	if err != nil {
		t.Fatalf("Post() error = %v", err)
	}
	defer resp.Body().Close()

	body, _ := io.ReadAll(resp.Body())
	if string(body) != "post data" {
		t.Errorf("Response body = %q, want %q", string(body), "post data")
	}
}

func TestRequest_Methods(t *testing.T) {
	h := NewHTTP()

	req, err := h.NewRequest("POST", "http://example.com/path", strings.NewReader("body"))
	if err != nil {
		t.Fatalf("NewRequest() error = %v", err)
	}

	realReq := req.RealRequest()

	// Test Method
	if realReq.Method != "POST" {
		t.Errorf("Method = %s, want POST", realReq.Method)
	}

	// Test URL
	if realReq.URL.String() != "http://example.com/path" {
		t.Errorf("URL = %s, want http://example.com/path", realReq.URL.String())
	}

	// Test Header
	realReq.Header.Set("X-Custom-Header", "value")
	if realReq.Header.Get("X-Custom-Header") != "value" {
		t.Errorf("Header = %q, want %q", realReq.Header.Get("X-Custom-Header"), "value")
	}

	// Test Body
	if realReq.Body == nil {
		t.Error("Body is nil")
	}
}

func TestResponse_Methods(t *testing.T) {
	server := httptest.NewServer(stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		w.Header().Set("X-Test", "value")
		w.WriteHeader(201)
		w.Write([]byte("response body"))
	}))
	defer server.Close()

	h := NewHTTP()
	resp, err := h.Get(server.URL)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	defer resp.Body().Close()

	// Test StatusCode
	if resp.StatusCode() != 201 {
		t.Errorf("StatusCode() = %d, want 201", resp.StatusCode())
	}

	// Test Status
	status := resp.Status()
	if !strings.Contains(status, "201") {
		t.Errorf("Status() = %q, want to contain '201'", status)
	}

	// Test Header
	if resp.Header().Get("X-Test") != "value" {
		t.Errorf("Header = %q, want %q", resp.Header().Get("X-Test"), "value")
	}

	// Test Body
	body, _ := io.ReadAll(resp.Body())
	if string(body) != "response body" {
		t.Errorf("Body = %q, want %q", string(body), "response body")
	}
}

func TestHTTP_CanonicalHeaderKey(t *testing.T) {
	h := NewHTTP()

	canonical := h.CanonicalHeaderKey("content-type")
	if canonical != "Content-Type" {
		t.Errorf("CanonicalHeaderKey() = %q, want %q", canonical, "Content-Type")
	}
}

func TestHTTP_StatusText(t *testing.T) {
	h := NewHTTP()

	text := h.StatusText(200)
	if text != "OK" {
		t.Errorf("StatusText(200) = %q, want %q", text, "OK")
	}

	text = h.StatusText(404)
	if text != "Not Found" {
		t.Errorf("StatusText(404) = %q, want %q", text, "Not Found")
	}
}
