package io

import (
	"io"
)

type PipeReader interface {
	Close() error
	CloseWithError(error) error
	Read([]byte) (int, error)
}

type pipeReaderFacade struct {
	realPipeReader *io.PipeReader
}

func (pr pipeReaderFacade) Close() error {
	return pr.realPipeReader.Close()
}

func (pr pipeReaderFacade) CloseWithError(err error) error {
	return pr.realPipeReader.CloseWithError(err)
}

func (pr pipeReaderFacade) Read(b []byte) (int, error) {
	return pr.realPipeReader.Read(b)
}
