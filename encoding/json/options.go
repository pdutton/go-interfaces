package json

import (
	"encoding/json"
)

type DecoderOption func(*json.Decoder)

func WithUseNumber() DecoderOption {
	return func(dec *json.Decoder) {
		dec.UseNumber()
	}
}

func WithDisallowUnknownFields() DecoderOption {
	return func(dec *json.Decoder) {
		dec.DisallowUnknownFields()
	}
}

type EncoderOption func(*json.Encoder)

func WithIndent(prefix, indent string) EncoderOption {
	return func(enc *json.Encoder) {
		enc.SetIndent(prefix, indent)
	}
}

func WithEscapeHTML(on bool) EncoderOption {
	return func(enc *json.Encoder) {
		enc.SetEscapeHTML(on)
	}
}
