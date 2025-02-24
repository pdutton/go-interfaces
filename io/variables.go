package io

import (
	"io"
)

var (
	EOF              = io.EOF
	ErrClosedPipe    = io.ErrClosedPipe
	ErrNoProgress    = io.ErrNoProgress
	ErrShortBuffer   = io.ErrShortBuffer
	ErrShortWrite    = io.ErrShortWrite
	ErrUnexpectedEOF = io.ErrUnexpectedEOF
)
