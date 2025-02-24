package io

import (
	"io"
)

type IO interface {
	// Functions:
	Copy(Writer, Reader) (int64, error)
	CopyBuffer(Writer, Reader, []byte) (int64, error)
	CopyN(Writer, Reader, int64) (int64, error)
	Pipe() (PipeReader, PipeWriter)
	ReadAll(Reader) ([]byte, error)
	ReadAtLeast(Reader, []byte, int) (int, error)
	ReadFull(Reader, []byte) (int, error)
	WriteString(Writer, string) (int, error)

	// Constructors:
	NewLimitedReader(Reader, int64) LimitedReader

	NewOffsetWriter(WriterAt, int64) OffsetWriter

	NopCloser(Reader) ReadCloser

	LimitReader(Reader, int64) Reader
	MultiReader(...Reader) Reader
	TeeReader(Reader, Writer) Reader

	NewSectionReader(ReaderAt, int64, int64) SectionReader

	MultiWriter(...Writer) Writer
}

type ioFacade struct {
}

func NewIO() ioFacade {
	return ioFacade{}
}

func (_ ioFacade) NopCloser(r Reader) ReadCloser {
	return io.NopCloser(r)
}

func (_ ioFacade) LimitReader(r Reader, n int64) Reader {
	return io.LimitReader(r, n)
}

func (_ ioFacade) MultiReader(readers ...Reader) Reader {
	return io.MultiReader(readers...)
}

func (_ ioFacade) TeeReader(r Reader, w Writer) Reader {
	return io.TeeReader(r, w)
}

func (_ ioFacade) MultiWriter(writers ...Writer) Writer {
	return io.MultiWriter(writers...)
}

func (_ ioFacade) Copy(w Writer, r Reader) (int64, error) {
	return io.Copy(w, r)
}

func (_ ioFacade) CopyBuffer(w Writer, r Reader, buf []byte) (int64, error) {
	return io.CopyBuffer(w, r, buf)
}

func (_ ioFacade) CopyN(w Writer, r Reader, nbr int64) (int64, error) {
	return io.CopyN(w, r, nbr)
}

func (_ ioFacade) Pipe() (PipeReader, PipeWriter) {
	return io.Pipe()
}

func (_ ioFacade) ReadAll(r Reader) ([]byte, error) {
	return io.ReadAll(r)
}

func (_ ioFacade) ReadAtLeast(r Reader, buf []byte, min int) (int, error) {
	return io.ReadAtLeast(r, buf, min)
}

func (_ ioFacade) ReadFull(r Reader, buf []byte) (int, error) {
	return io.ReadFull(r, buf)
}

func (_ ioFacade) WriteString(w Writer, str string) (int, error) {
	return io.WriteString(w, str)
}
