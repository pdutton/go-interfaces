// Package json provides an interface to functions and types
// in the standard encoding/json package to facilitate mocking.
package json

import (
	"bytes"
	"encoding/json"
	"io"
)

type JSON interface {
	// Functions:
	Marshal(v any) ([]byte, error)
	MarshalIndent(v any, prefix, indent string) ([]byte, error)
	Unmarshal(data []byte, v any) error
	Compact(dst *bytes.Buffer, src []byte) error
	HTMLEscape(dst *bytes.Buffer, src []byte)
	Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
	Valid(data []byte) bool

	// Constructors:
	NewDecoder(r io.Reader, options ...DecoderOption) Decoder
	NewEncoder(w io.Writer, options ...EncoderOption) Encoder
}

type jsonFacade struct {
}

func NewJSON() JSON {
	return jsonFacade{}
}

func (_ jsonFacade) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (_ jsonFacade) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (_ jsonFacade) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (_ jsonFacade) Compact(dst *bytes.Buffer, src []byte) error {
	return json.Compact(dst, src)
}

func (_ jsonFacade) HTMLEscape(dst *bytes.Buffer, src []byte) {
	json.HTMLEscape(dst, src)
}

func (_ jsonFacade) Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}

func (_ jsonFacade) Valid(data []byte) bool {
	return json.Valid(data)
}

// Marshal returns the JSON encoding of v by calling encoding/json.Marshal.
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalIndent returns formatted JSON encoding of v by calling encoding/json.MarshalIndent.
func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal parses the JSON-encoded data and stores the result by calling encoding/json.Unmarshal.
func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
