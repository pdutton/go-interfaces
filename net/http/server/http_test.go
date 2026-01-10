package http

import (
	"bytes"
	"io"
	stdhttp "net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNewHTTP(t *testing.T) {
	h := NewHTTP()
	_ = h
}

func TestHTTP_HandleFunc(t *testing.T) {
	// Create a test server with HandleFunc
	mux := stdhttp.NewServeMux()
	called := false

	handler := func(w ResponseWriter, r *stdhttp.Request) {
		called = true
		w.Write([]byte("handler called"))
	}

	// Register handler
	mux.HandleFunc("/test", func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		handler(w, r)
	})

	// Create test server
	server := httptest.NewServer(mux)
	defer server.Close()

	// Make request
	resp, err := stdhttp.Get(server.URL + "/test")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	defer resp.Body.Close()

	if !called {
		t.Error("Handler was not called")
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "handler called" {
		t.Errorf("Response = %q, want %q", string(body), "handler called")
	}
}

func TestHTTP_Handle(t *testing.T) {
	mux := stdhttp.NewServeMux()
	handler := stdhttp.HandlerFunc(func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		w.Write([]byte("handle"))
	})

	mux.Handle("/test", handler)

	server := httptest.NewServer(mux)
	defer server.Close()

	resp, err := stdhttp.Get(server.URL + "/test")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "handle" {
		t.Errorf("Response = %q, want %q", string(body), "handle")
	}
}

func TestHTTP_Error(t *testing.T) {
	h := NewHTTP()

	rec := httptest.NewRecorder()
	h.Error(rec, "test error", 500)

	if rec.Code != 500 {
		t.Errorf("Status code = %d, want 500", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "test error") {
		t.Errorf("Body = %q, want to contain 'test error'", body)
	}
}

func TestHTTP_NotFound(t *testing.T) {
	h := NewHTTP()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/notfound", nil)

	h.NotFound(rec, req)

	if rec.Code != 404 {
		t.Errorf("Status code = %d, want 404", rec.Code)
	}
}

func TestHTTP_Redirect(t *testing.T) {
	h := NewHTTP()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/old", nil)

	h.Redirect(rec, req, "/new", 302)

	if rec.Code != 302 {
		t.Errorf("Status code = %d, want 302", rec.Code)
	}

	location := rec.Header().Get("Location")
	if location != "/new" {
		t.Errorf("Location = %q, want %q", location, "/new")
	}
}

func TestHTTP_ServeContent(t *testing.T) {
	h := NewHTTP()

	content := []byte("test content")
	reader := bytes.NewReader(content)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/file.txt", nil)

	modTime := time.Now()
	h.ServeContent(rec, req, "file.txt", modTime, reader)

	if rec.Code != 200 {
		t.Errorf("Status code = %d, want 200", rec.Code)
	}

	body := rec.Body.String()
	if body != "test content" {
		t.Errorf("Body = %q, want %q", body, "test content")
	}
}

func TestHTTP_DetectContentType(t *testing.T) {
	h := NewHTTP()

	tests := []struct {
		name     string
		data     []byte
		expected string
	}{
		{"HTML", []byte("<html>"), "text/html; charset=utf-8"},
		{"JSON", []byte(`{"key":"value"}`), "text/plain; charset=utf-8"},
		{"Binary", []byte{0xFF, 0xD8, 0xFF}, "image/jpeg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contentType := h.DetectContentType(tt.data)
			if contentType != tt.expected {
				t.Errorf("DetectContentType() = %q, want %q", contentType, tt.expected)
			}
		})
	}
}

func TestHTTP_StatusText(t *testing.T) {
	h := NewHTTP()

	tests := []struct {
		code int
		text string
	}{
		{200, "OK"},
		{404, "Not Found"},
		{500, "Internal Server Error"},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			text := h.StatusText(tt.code)
			if text != tt.text {
				t.Errorf("StatusText(%d) = %q, want %q", tt.code, text, tt.text)
			}
		})
	}
}

func TestHTTP_CanonicalHeaderKey(t *testing.T) {
	h := NewHTTP()

	tests := []struct {
		input    string
		expected string
	}{
		{"content-type", "Content-Type"},
		{"x-custom-header", "X-Custom-Header"},
		{"ACCEPT", "Accept"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			canonical := h.CanonicalHeaderKey(tt.input)
			if canonical != tt.expected {
				t.Errorf("CanonicalHeaderKey(%q) = %q, want %q", tt.input, canonical, tt.expected)
			}
		})
	}
}

func TestHTTP_SetCookie(t *testing.T) {
	h := NewHTTP()

	rec := httptest.NewRecorder()
	cookie := &Cookie{
		Name:  "test",
		Value: "value",
	}

	h.SetCookie(rec, cookie)

	cookies := rec.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("Got %d cookies, want 1", len(cookies))
	}

	if cookies[0].Name != "test" {
		t.Errorf("Cookie name = %q, want %q", cookies[0].Name, "test")
	}

	if cookies[0].Value != "value" {
		t.Errorf("Cookie value = %q, want %q", cookies[0].Value, "value")
	}
}

func TestHTTP_ParseHTTPVersion(t *testing.T) {
	h := NewHTTP()

	major, minor, ok := h.ParseHTTPVersion("HTTP/1.1")
	if !ok {
		t.Fatal("ParseHTTPVersion() ok = false, want true")
	}

	if major != 1 {
		t.Errorf("Major = %d, want 1", major)
	}

	if minor != 1 {
		t.Errorf("Minor = %d, want 1", minor)
	}
}

func TestHTTP_ParseTime(t *testing.T) {
	h := NewHTTP()

	// RFC1123 format
	timeStr := "Mon, 02 Jan 2006 15:04:05 GMT"
	parsed, err := h.ParseTime(timeStr)
	if err != nil {
		t.Errorf("ParseTime() error = %v", err)
		return
	}

	if parsed.IsZero() {
		t.Error("ParseTime() returned zero time")
	}
}

func TestHTTP_MaxBytesReader(t *testing.T) {
	h := NewHTTP()

	rec := httptest.NewRecorder()
	body := io.NopCloser(strings.NewReader("test data longer than limit"))

	limited := h.MaxBytesReader(rec, body, 5)
	defer limited.Close()

	data, err := io.ReadAll(limited)
	if err == nil {
		t.Error("Expected error when reading beyond limit")
	}

	if len(data) > 5 {
		t.Errorf("Read %d bytes, want max 5", len(data))
	}
}

func TestResponseWriter_Methods(t *testing.T) {
	rec := httptest.NewRecorder()

	// Test Header
	rec.Header().Set("X-Test", "value")
	if rec.Header().Get("X-Test") != "value" {
		t.Error("Header not set correctly")
	}

	// Test WriteHeader (must be called before Write)
	rec.WriteHeader(201)
	if rec.Code != 201 {
		t.Errorf("WriteHeader() code = %d, want 201", rec.Code)
	}

	// Test Write
	n, err := rec.Write([]byte("test"))
	if err != nil {
		t.Errorf("Write() error = %v", err)
	}
	if n != 4 {
		t.Errorf("Write() = %d, want 4", n)
	}
}

func TestHTTP_ServeFile(t *testing.T) {
	h := NewHTTP()

	// Create a temp file
	tmpDir := t.TempDir()
	tmpFile := tmpDir + "/test.txt"

	// Write some content
	content := []byte("file content")
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test.txt", nil)

	// Write the file using os package
	if err := os.WriteFile(tmpFile, content, 0644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	// Serve it
	h.ServeFile(rec, req, tmpFile)

	if rec.Code != 200 {
		t.Errorf("Status code = %d, want 200", rec.Code)
	}

	body := rec.Body.String()
	if body != "file content" {
		t.Errorf("Body = %q, want %q", body, "file content")
	}
}
