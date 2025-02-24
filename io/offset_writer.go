package io

import (
	"io"
)

type OffsetWriter interface {
	Seek(int64, int) (int64, error)
	Write([]byte) (int, error)
	WriteAt([]byte, int64) (int, error)
}

type offsetWriterFacade struct {
	realOffsetWriter *io.OffsetWriter
}

func (_ ioFacade) NewOffsetWriter(w WriterAt, off int64) OffsetWriter {
	return offsetWriterFacade{
		realOffsetWriter: io.NewOffsetWriter(w, off),
	}
}

func (ow offsetWriterFacade) Seek(offset int64, whence int) (int64, error) {
	return ow.realOffsetWriter.Seek(offset, whence)
}

func (ow offsetWriterFacade) Write(p []byte) (int, error) {
	return ow.realOffsetWriter.Write(p)
}

func (ow offsetWriterFacade) WriteAt(p []byte, off int64) (int, error) {
	return ow.realOffsetWriter.WriteAt(p, off)
}
