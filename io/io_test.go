package io

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewIO(t *testing.T) {
	io := NewIO()
	_ = io
}

func TestIO_Copy(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")
	var dst bytes.Buffer

	n, err := io.Copy(&dst, src)
	if err != nil {
		t.Errorf("Copy() error = %v", err)
	}
	if n != 11 {
		t.Errorf("Copy() copied %d bytes, want 11", n)
	}
	if dst.String() != "hello world" {
		t.Errorf("Copy() result = %q, want %q", dst.String(), "hello world")
	}
}

func TestIO_CopyBuffer(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")
	var dst bytes.Buffer
	buf := make([]byte, 5)

	n, err := io.CopyBuffer(&dst, src, buf)
	if err != nil {
		t.Errorf("CopyBuffer() error = %v", err)
	}
	if n != 11 {
		t.Errorf("CopyBuffer() copied %d bytes, want 11", n)
	}
	if dst.String() != "hello world" {
		t.Errorf("CopyBuffer() result = %q, want %q", dst.String(), "hello world")
	}
}

func TestIO_CopyN(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")
	var dst bytes.Buffer

	n, err := io.CopyN(&dst, src, 5)
	if err != nil {
		t.Errorf("CopyN() error = %v", err)
	}
	if n != 5 {
		t.Errorf("CopyN() copied %d bytes, want 5", n)
	}
	if dst.String() != "hello" {
		t.Errorf("CopyN() result = %q, want %q", dst.String(), "hello")
	}
}

func TestIO_ReadAll(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")

	data, err := io.ReadAll(src)
	if err != nil {
		t.Errorf("ReadAll() error = %v", err)
	}
	if string(data) != "hello world" {
		t.Errorf("ReadAll() = %q, want %q", string(data), "hello world")
	}
}

func TestIO_ReadAtLeast(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")
	buf := make([]byte, 20)

	n, err := io.ReadAtLeast(src, buf, 5)
	if err != nil {
		t.Errorf("ReadAtLeast() error = %v", err)
	}
	if n < 5 {
		t.Errorf("ReadAtLeast() read %d bytes, want at least 5", n)
	}
}

func TestIO_ReadFull(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")
	buf := make([]byte, 5)

	n, err := io.ReadFull(src, buf)
	if err != nil {
		t.Errorf("ReadFull() error = %v", err)
	}
	if n != 5 {
		t.Errorf("ReadFull() read %d bytes, want 5", n)
	}
	if string(buf) != "hello" {
		t.Errorf("ReadFull() = %q, want %q", string(buf), "hello")
	}
}

func TestIO_WriteString(t *testing.T) {
	io := NewIO()

	var buf bytes.Buffer

	n, err := io.WriteString(&buf, "hello world")
	if err != nil {
		t.Errorf("WriteString() error = %v", err)
	}
	if n != 11 {
		t.Errorf("WriteString() wrote %d bytes, want 11", n)
	}
	if buf.String() != "hello world" {
		t.Errorf("WriteString() result = %q, want %q", buf.String(), "hello world")
	}
}

func TestIO_Pipe(t *testing.T) {
	io := NewIO()

	pr, pw := io.Pipe()
	if pr == nil || pw == nil {
		t.Fatal("Pipe() returned nil")
	}

	// Write in a goroutine
	go func() {
		pw.Write([]byte("test data"))
		pw.Close()
	}()

	// Read from the pipe
	buf := make([]byte, 20)
	n, err := pr.Read(buf)
	if err != nil && err.Error() != "EOF" {
		t.Errorf("Read() error = %v", err)
	}
	if string(buf[:n]) != "test data" {
		t.Errorf("Read() = %q, want %q", string(buf[:n]), "test data")
	}
}

func TestIO_LimitReader(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")
	limited := io.LimitReader(src, 5)

	buf := make([]byte, 20)
	n, _ := limited.Read(buf)
	if n != 5 {
		t.Errorf("LimitReader read %d bytes, want 5", n)
	}
	if string(buf[:n]) != "hello" {
		t.Errorf("LimitReader = %q, want %q", string(buf[:n]), "hello")
	}
}

func TestIO_MultiReader(t *testing.T) {
	io := NewIO()

	r1 := strings.NewReader("hello ")
	r2 := strings.NewReader("world")

	mr := io.MultiReader(r1, r2)

	// Read all data - may require multiple reads
	var buf bytes.Buffer
	buf.ReadFrom(mr)
	if buf.String() != "hello world" {
		t.Errorf("MultiReader = %q, want %q", buf.String(), "hello world")
	}
}

func TestIO_TeeReader(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")
	var tee bytes.Buffer

	tr := io.TeeReader(src, &tee)

	buf := make([]byte, 20)
	n, _ := tr.Read(buf)
	if string(buf[:n]) != "hello world" {
		t.Errorf("TeeReader read = %q, want %q", string(buf[:n]), "hello world")
	}
	if tee.String() != "hello world" {
		t.Errorf("TeeReader wrote = %q, want %q", tee.String(), "hello world")
	}
}

func TestIO_MultiWriter(t *testing.T) {
	io := NewIO()

	var w1, w2 bytes.Buffer

	mw := io.MultiWriter(&w1, &w2)

	mw.Write([]byte("hello world"))

	if w1.String() != "hello world" {
		t.Errorf("MultiWriter w1 = %q, want %q", w1.String(), "hello world")
	}
	if w2.String() != "hello world" {
		t.Errorf("MultiWriter w2 = %q, want %q", w2.String(), "hello world")
	}
}

func TestIO_NopCloser(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")
	rc := io.NopCloser(src)

	// Read from it
	buf := make([]byte, 20)
	n, _ := rc.Read(buf)
	if string(buf[:n]) != "hello world" {
		t.Errorf("NopCloser read = %q, want %q", string(buf[:n]), "hello world")
	}

	// Close should not error
	if err := rc.Close(); err != nil {
		t.Errorf("NopCloser Close() error = %v", err)
	}
}

func TestIO_NewLimitedReader(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")
	lr := io.NewLimitedReader(src, 5)

	buf := make([]byte, 20)
	n, _ := lr.Read(buf)
	if n != 5 {
		t.Errorf("LimitedReader read %d bytes, want 5", n)
	}
	if string(buf[:n]) != "hello" {
		t.Errorf("LimitedReader = %q, want %q", string(buf[:n]), "hello")
	}
}

func TestIO_NewOffsetWriter(t *testing.T) {
	io := NewIO()

	// Use a byte slice that we can write to
	data := make([]byte, 20)
	wa := bytesSliceWriterAt{data}

	ow := io.NewOffsetWriter(wa, 5)

	// Write should happen at offset 5
	n, err := ow.Write([]byte("test"))
	if err != nil {
		t.Errorf("Write() error = %v", err)
	}
	if n != 4 {
		t.Errorf("Write() = %d, want 4", n)
	}

	// Check that it wrote at offset 5
	if string(data[5:9]) != "test" {
		t.Errorf("OffsetWriter wrote %q at offset 5, want %q", string(data[5:9]), "test")
	}
}

func TestIO_NewSectionReader(t *testing.T) {
	io := NewIO()

	data := []byte("hello world")
	ra := bytes.NewReader(data)

	sr := io.NewSectionReader(ra, 6, 5)

	buf := make([]byte, 10)
	n, _ := sr.Read(buf)
	if string(buf[:n]) != "world" {
		t.Errorf("SectionReader = %q, want %q", string(buf[:n]), "world")
	}

	// Test Size
	if sr.Size() != 5 {
		t.Errorf("SectionReader Size() = %d, want 5", sr.Size())
	}
}

// Helper type for OffsetWriter test
type bytesSliceWriterAt struct {
	data []byte
}

func (b bytesSliceWriterAt) WriteAt(p []byte, off int64) (n int, err error) {
	return copy(b.data[off:], p), nil
}

func TestPipeReader(t *testing.T) {
	io := NewIO()

	pr, pw := io.Pipe()

	// Test PipeReader methods
	go func() {
		pw.Write([]byte("test"))
		pw.Close()
	}()

	buf := make([]byte, 10)
	n, _ := pr.Read(buf)
	if string(buf[:n]) != "test" {
		t.Errorf("PipeReader Read() = %q, want %q", string(buf[:n]), "test")
	}

	if err := pr.Close(); err != nil {
		t.Errorf("PipeReader Close() error = %v", err)
	}
}

func TestPipeReader_CloseWithError(t *testing.T) {
	io := NewIO()

	pr, pw := io.Pipe()

	go func() {
		pw.Write([]byte("test"))
		pr.CloseWithError(nil)
	}()

	buf := make([]byte, 10)
	pr.Read(buf)
	pw.Close()
}

func TestPipeWriter(t *testing.T) {
	io := NewIO()

	pr, pw := io.Pipe()

	go func() {
		buf := make([]byte, 10)
		pr.Read(buf)
		pr.Close()
	}()

	// Test PipeWriter methods
	n, err := pw.Write([]byte("test"))
	if err != nil {
		t.Errorf("PipeWriter Write() error = %v", err)
	}
	if n != 4 {
		t.Errorf("PipeWriter Write() = %d, want 4", n)
	}

	if err := pw.Close(); err != nil {
		t.Errorf("PipeWriter Close() error = %v", err)
	}
}

func TestSectionReader_Methods(t *testing.T) {
	io := NewIO()

	data := []byte("hello world")
	ra := bytes.NewReader(data)

	sr := io.NewSectionReader(ra, 0, 11)

	// Test ReadAt
	buf := make([]byte, 5)
	n, err := sr.ReadAt(buf, 6)
	if err != nil {
		t.Errorf("ReadAt() error = %v", err)
	}
	if string(buf[:n]) != "world" {
		t.Errorf("ReadAt() = %q, want %q", string(buf[:n]), "world")
	}

	// Test Seek
	pos, err := sr.Seek(6, 0)
	if err != nil {
		t.Errorf("Seek() error = %v", err)
	}
	if pos != 6 {
		t.Errorf("Seek() = %d, want 6", pos)
	}

	// Test Outer
	outer, off, size := sr.Outer()
	if outer == nil {
		t.Error("Outer() returned nil")
	}
	if off != 0 {
		t.Errorf("Outer() offset = %d, want 0", off)
	}
	if size != 11 {
		t.Errorf("Outer() size = %d, want 11", size)
	}
}

func TestLimitedReader_Methods(t *testing.T) {
	io := NewIO()

	src := strings.NewReader("hello world")
	lr := io.NewLimitedReader(src, 5)

	// Test Read
	buf := make([]byte, 10)
	n, _ := lr.Read(buf)
	if n != 5 {
		t.Errorf("Read() = %d, want 5", n)
	}
}

func TestOffsetWriter_Methods(t *testing.T) {
	io := NewIO()

	data := make([]byte, 20)
	wa := bytesSliceWriterAt{data}

	ow := io.NewOffsetWriter(wa, 0)

	// Test WriteAt
	n, err := ow.WriteAt([]byte("test"), 0)
	if err != nil {
		t.Errorf("WriteAt() error = %v", err)
	}
	if n != 4 {
		t.Errorf("WriteAt() = %d, want 4", n)
	}
	if string(data[0:4]) != "test" {
		t.Errorf("WriteAt() wrote %q, want %q", string(data[0:4]), "test")
	}

	// Test Seek
	pos, err := ow.Seek(5, 0)
	if err != nil {
		t.Errorf("Seek() error = %v", err)
	}
	if pos != 5 {
		t.Errorf("Seek() = %d, want 5", pos)
	}
}
