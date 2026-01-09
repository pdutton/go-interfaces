package json

import (
	"encoding/json"
	"io"
)

type Encoder interface {
	Encode(v any) error
	SetIndent(prefix, indent string)
	SetEscapeHTML(on bool)
	Nub() *json.Encoder
}

type encoderFacade struct {
	realEncoder *json.Encoder
}

func (_ jsonFacade) NewEncoder(w io.Writer, options ...EncoderOption) Encoder {
	enc := json.NewEncoder(w)

	for _, opt := range options {
		opt(enc)
	}

	return encoderFacade{realEncoder: enc}
}

func WrapEncoder(enc *json.Encoder) Encoder {
	return encoderFacade{realEncoder: enc}
}

func (e encoderFacade) Encode(v any) error {
	return e.realEncoder.Encode(v)
}

func (e encoderFacade) SetIndent(prefix, indent string) {
	e.realEncoder.SetIndent(prefix, indent)
}

func (e encoderFacade) SetEscapeHTML(on bool) {
	e.realEncoder.SetEscapeHTML(on)
}

func (e encoderFacade) Nub() *json.Encoder {
	return e.realEncoder
}
