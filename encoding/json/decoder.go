package json

import (
	"encoding/json"
	"io"
)

type Decoder interface {
	Decode(v any) error
	Buffered() io.Reader
	InputOffset() int64
	More() bool
	Token() (Token, error)
	Nub() *json.Decoder
}

type decoderFacade struct {
	realDecoder *json.Decoder
}

func (_ jsonFacade) NewDecoder(r io.Reader, options ...DecoderOption) Decoder {
	return NewDecoder(r, options...)
}

func WrapDecoder(dec *json.Decoder) Decoder {
	return decoderFacade{realDecoder: dec}
}

// NewDecoder creates a new Decoder that reads from r with optional configuration.
func NewDecoder(r io.Reader, options ...DecoderOption) Decoder {
	dec := json.NewDecoder(r)

	for _, opt := range options {
		opt(dec)
	}

	return decoderFacade{realDecoder: dec}
}

func (d decoderFacade) Decode(v any) error {
	return d.realDecoder.Decode(v)
}

func (d decoderFacade) Buffered() io.Reader {
	return d.realDecoder.Buffered()
}

func (d decoderFacade) InputOffset() int64 {
	return d.realDecoder.InputOffset()
}

func (d decoderFacade) More() bool {
	return d.realDecoder.More()
}

func (d decoderFacade) Token() (Token, error) {
	return d.realDecoder.Token()
}

func (d decoderFacade) Nub() *json.Decoder {
	return d.realDecoder
}
