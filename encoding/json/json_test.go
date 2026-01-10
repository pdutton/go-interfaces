package json

import (
	"bytes"
	stdJson "encoding/json"
	"strings"
	"testing"
)

type testStruct struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestNewJSON(t *testing.T) {
	j := NewJSON()
	_ = j
}

func TestJSON_Marshal(t *testing.T) {
	j := NewJSON()

	tests := []struct {
		name    string
		input   any
		wantErr bool
	}{
		{"struct", testStruct{Name: "test", Value: 42}, false},
		{"map", map[string]string{"key": "value"}, false},
		{"slice", []int{1, 2, 3}, false},
		{"string", "hello", false},
		{"number", 42, false},
		{"bool", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := j.Marshal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(data) == 0 {
				t.Error("Marshal() returned empty data")
			}
		})
	}
}

func TestJSON_MarshalIndent(t *testing.T) {
	j := NewJSON()

	data, err := j.MarshalIndent(testStruct{Name: "test", Value: 42}, "", "  ")
	if err != nil {
		t.Errorf("MarshalIndent() error = %v", err)
		return
	}

	// Check that data is properly indented
	if !strings.Contains(string(data), "\n") {
		t.Error("MarshalIndent() didn't produce indented output")
	}
}

func TestJSON_Unmarshal(t *testing.T) {
	j := NewJSON()

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid JSON", `{"name":"test","value":42}`, false},
		{"invalid JSON", `{invalid}`, true},
		{"empty", ``, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result testStruct
			err := j.Unmarshal([]byte(tt.input), &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && result.Name != "test" {
				t.Errorf("Unmarshal() didn't populate struct correctly")
			}
		})
	}
}

func TestJSON_Valid(t *testing.T) {
	j := NewJSON()

	tests := []struct {
		name  string
		input string
		valid bool
	}{
		{"valid object", `{"key":"value"}`, true},
		{"valid array", `[1,2,3]`, true},
		{"valid string", `"hello"`, true},
		{"valid number", `42`, true},
		{"valid bool", `true`, true},
		{"invalid", `{invalid}`, false},
		{"empty", ``, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := j.Valid([]byte(tt.input))
			if result != tt.valid {
				t.Errorf("Valid(%q) = %v, want %v", tt.input, result, tt.valid)
			}
		})
	}
}

func TestJSON_Compact(t *testing.T) {
	j := NewJSON()

	input := []byte(`{
		"name": "test",
		"value": 42
	}`)
	var buf bytes.Buffer

	err := j.Compact(&buf, input)
	if err != nil {
		t.Errorf("Compact() error = %v", err)
		return
	}

	result := buf.String()
	if strings.Contains(result, "\n") || strings.Contains(result, "\t") {
		t.Error("Compact() didn't remove whitespace")
	}
}

func TestJSON_HTMLEscape(t *testing.T) {
	j := NewJSON()

	input := []byte(`{"html":"<script>alert('xss')</script>"}`)
	var buf bytes.Buffer

	j.HTMLEscape(&buf, input)

	result := buf.String()
	if strings.Contains(result, "<script>") {
		t.Error("HTMLEscape() didn't escape HTML")
	}
}

func TestJSON_Indent(t *testing.T) {
	j := NewJSON()

	input := []byte(`{"name":"test","value":42}`)
	var buf bytes.Buffer

	err := j.Indent(&buf, input, "", "  ")
	if err != nil {
		t.Errorf("Indent() error = %v", err)
		return
	}

	result := buf.String()
	if !strings.Contains(result, "\n") {
		t.Error("Indent() didn't add newlines")
	}
}

func TestNewDecoder(t *testing.T) {
	j := NewJSON()

	input := strings.NewReader(`{"name":"test","value":42}`)
	dec := j.NewDecoder(input)
	_ = dec
}

func TestDecoder_Decode(t *testing.T) {
	j := NewJSON()

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"valid JSON", `{"name":"test","value":42}`, false},
		{"invalid JSON", `{invalid}`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec := j.NewDecoder(strings.NewReader(tt.input))
			var result testStruct
			err := dec.Decode(&result)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && result.Name != "test" {
				t.Error("Decode() didn't populate struct correctly")
			}
		})
	}
}

func TestDecoder_More(t *testing.T) {
	j := NewJSON()

	input := strings.NewReader(`{"name":"test"}{"name":"test2"}`)
	dec := j.NewDecoder(input)

	// Decode first object
	var first testStruct
	if err := dec.Decode(&first); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	// Check if more data is available
	if !dec.More() {
		t.Error("More() = false, want true (second object available)")
	}

	// Decode second object
	var second testStruct
	if err := dec.Decode(&second); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	// No more data should be available
	if dec.More() {
		t.Error("More() = true, want false (no more data)")
	}
}

func TestDecoder_Token(t *testing.T) {
	j := NewJSON()

	input := strings.NewReader(`{"name":"test"}`)
	dec := j.NewDecoder(input)

	// First token should be start of object
	token, err := dec.Token()
	if err != nil {
		t.Fatalf("Token() error = %v", err)
	}
	if token == nil {
		t.Error("Token() returned nil")
	}
}

func TestDecoder_Buffered(t *testing.T) {
	j := NewJSON()

	input := strings.NewReader(`{"name":"test"}extra data`)
	dec := j.NewDecoder(input)

	var result testStruct
	if err := dec.Decode(&result); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	// Get buffered data
	buffered := dec.Buffered()
	if buffered == nil {
		t.Error("Buffered() returned nil")
	}
}

func TestDecoder_InputOffset(t *testing.T) {
	j := NewJSON()

	input := strings.NewReader(`{"name":"test"}`)
	dec := j.NewDecoder(input)

	var result testStruct
	if err := dec.Decode(&result); err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	// Check that offset has advanced
	offset := dec.InputOffset()
	if offset == 0 {
		t.Error("InputOffset() = 0, want > 0 after decoding")
	}
}

func TestNewEncoder(t *testing.T) {
	j := NewJSON()

	var buf bytes.Buffer
	enc := j.NewEncoder(&buf)
	_ = enc
}

func TestEncoder_Encode(t *testing.T) {
	j := NewJSON()

	tests := []struct {
		name    string
		input   any
		wantErr bool
	}{
		{"struct", testStruct{Name: "test", Value: 42}, false},
		{"map", map[string]string{"key": "value"}, false},
		{"slice", []int{1, 2, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			enc := j.NewEncoder(&buf)
			err := enc.Encode(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && buf.Len() == 0 {
				t.Error("Encode() didn't write any data")
			}
		})
	}
}

func TestEncoder_SetIndent(t *testing.T) {
	j := NewJSON()

	var buf bytes.Buffer
	enc := j.NewEncoder(&buf)
	enc.SetIndent("", "  ")

	if err := enc.Encode(testStruct{Name: "test", Value: 42}); err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	result := buf.String()
	if !strings.Contains(result, "\n") {
		t.Error("SetIndent() didn't produce indented output")
	}
}

func TestEncoder_SetEscapeHTML(t *testing.T) {
	j := NewJSON()

	var buf bytes.Buffer
	enc := j.NewEncoder(&buf)
	enc.SetEscapeHTML(false)

	input := map[string]string{"html": "<tag>"}
	if err := enc.Encode(input); err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	result := buf.String()
	if !strings.Contains(result, "<tag>") {
		t.Error("SetEscapeHTML(false) still escaped HTML")
	}
}

func TestPackageLevelFunctions(t *testing.T) {
	// Test package-level Marshal
	data, err := Marshal(testStruct{Name: "test", Value: 42})
	if err != nil {
		t.Errorf("Marshal() error = %v", err)
	}
	if len(data) == 0 {
		t.Error("Marshal() returned empty data")
	}

	// Test package-level MarshalIndent
	indented, err := MarshalIndent(testStruct{Name: "test", Value: 42}, "", "  ")
	if err != nil {
		t.Errorf("MarshalIndent() error = %v", err)
	}
	if len(indented) == 0 {
		t.Error("MarshalIndent() returned empty data")
	}

	// Test package-level Unmarshal
	var result testStruct
	if err := Unmarshal(data, &result); err != nil {
		t.Errorf("Unmarshal() error = %v", err)
	}
	if result.Name != "test" {
		t.Error("Unmarshal() didn't populate struct correctly")
	}
}

func TestDecoderOptions(t *testing.T) {
	j := NewJSON()

	// Test WithUseNumber
	t.Run("WithUseNumber", func(t *testing.T) {
		input := strings.NewReader(`{"value":42.5}`)
		dec := j.NewDecoder(input, WithUseNumber())

		var result map[string]any
		if err := dec.Decode(&result); err != nil {
			t.Fatalf("Decode() error = %v", err)
		}
	})

	// Test WithDisallowUnknownFields
	t.Run("WithDisallowUnknownFields", func(t *testing.T) {
		input := strings.NewReader(`{"name":"test","value":42,"unknown":"field"}`)
		dec := j.NewDecoder(input, WithDisallowUnknownFields())

		var result testStruct
		err := dec.Decode(&result)
		// Should error on unknown field
		if err == nil {
			t.Error("Expected error for unknown field, got nil")
		}
	})
}

func TestEncoderOptions(t *testing.T) {
	j := NewJSON()

	// Test WithIndent
	t.Run("WithIndent", func(t *testing.T) {
		var buf bytes.Buffer
		enc := j.NewEncoder(&buf, WithIndent("", "  "))

		if err := enc.Encode(testStruct{Name: "test", Value: 42}); err != nil {
			t.Fatalf("Encode() error = %v", err)
		}

		result := buf.String()
		if !strings.Contains(result, "\n") {
			t.Error("WithIndent() didn't produce indented output")
		}
	})

	// Test WithEscapeHTML
	t.Run("WithEscapeHTML", func(t *testing.T) {
		var buf bytes.Buffer
		enc := j.NewEncoder(&buf, WithEscapeHTML(false))

		input := map[string]string{"html": "<tag>"}
		if err := enc.Encode(input); err != nil {
			t.Fatalf("Encode() error = %v", err)
		}

		result := buf.String()
		if !strings.Contains(result, "<tag>") {
			t.Error("WithEscapeHTML(false) still escaped HTML")
		}
	})
}

func TestWrapDecoder(t *testing.T) {
	input := strings.NewReader(`{"name":"test","value":42}`)
	stdDecoder := stdJson.NewDecoder(input)

	wrapped := WrapDecoder(stdDecoder)
	var result testStruct
	if err := wrapped.Decode(&result); err != nil {
		t.Errorf("WrapDecoder Decode() error = %v", err)
	}
	if result.Name != "test" {
		t.Error("WrapDecoder didn't work correctly")
	}
}

func TestWrapEncoder(t *testing.T) {
	var buf bytes.Buffer
	stdEncoder := stdJson.NewEncoder(&buf)

	wrapped := WrapEncoder(stdEncoder)
	if err := wrapped.Encode(testStruct{Name: "test", Value: 42}); err != nil {
		t.Errorf("WrapEncoder Encode() error = %v", err)
	}
	if buf.Len() == 0 {
		t.Error("WrapEncoder didn't write data")
	}
}

func TestDecoder_Nub(t *testing.T) {
	input := strings.NewReader(`{"name":"test"}`)
	dec := NewDecoder(input)

	nub := dec.Nub()
	if nub == nil {
		t.Error("Nub() returned nil")
	}
}

func TestEncoder_Nub(t *testing.T) {
	var buf bytes.Buffer
	enc := NewEncoder(&buf)

	nub := enc.Nub()
	if nub == nil {
		t.Error("Nub() returned nil")
	}
}
