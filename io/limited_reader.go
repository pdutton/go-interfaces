package io

import (
	"io"
)

type LimitedReader interface {
	Read([]byte) (int, error)
}

type limitedReaderFacade struct {
	realLimitedReader *io.LimitedReader
}

func (_ ioFacade) NewLimitedReader(r Reader, n int64) LimitedReader {
	return limitedReaderFacade{
		realLimitedReader: &io.LimitedReader{
			R: r,
			N: n,
		},
	}
}

func (lr limitedReaderFacade) Read(p []byte) (int, error) {
	return lr.realLimitedReader.Read(p)
}
