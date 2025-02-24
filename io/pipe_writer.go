package io

import (
	"io"
)

type PipeWriter interface {
	Close() error
	CloseWithError(error) error
	Write([]byte) (int, error)
}

type pipeWriterFacade struct {
	realPipeWriter *io.PipeWriter
}

func (pw pipeWriterFacade) Close() error {
	return pw.realPipeWriter.Close()
}

func (pw pipeWriterFacade) CloseWithError(err error) error {
	return pw.realPipeWriter.CloseWithError(err)
}

func (pw pipeWriterFacade) Write(data []byte) (int, error) {
	return pw.realPipeWriter.Write(data)
}
