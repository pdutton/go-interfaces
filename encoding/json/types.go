package json

import (
	"encoding/json"
)

// Value types
type Delim = json.Delim
type Number = json.Number
type RawMessage = json.RawMessage
type Token = json.Token

// Interfaces for custom marshaling/unmarshaling
type Marshaler = json.Marshaler
type Unmarshaler = json.Unmarshaler

// Error types
type InvalidUnmarshalError = json.InvalidUnmarshalError
type MarshalerError = json.MarshalerError
type SyntaxError = json.SyntaxError
type UnmarshalTypeError = json.UnmarshalTypeError
type UnsupportedTypeError = json.UnsupportedTypeError
type UnsupportedValueError = json.UnsupportedValueError
