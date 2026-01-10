package path

import (
	"testing"
)

func TestNewPath(t *testing.T) {
	p := NewPath()
	_ = p
}

func TestPath_Base(t *testing.T) {
	p := NewPath()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple file", "/path/to/file.txt", "file.txt"},
		{"directory", "/path/to/dir/", "dir"},
		{"root", "/", "/"},
		{"empty", "", "."},
		{"dot", ".", "."},
		{"just filename", "file.txt", "file.txt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.Base(tt.input)
			if result != tt.expected {
				t.Errorf("Base(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPath_Clean(t *testing.T) {
	p := NewPath()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple path", "/path/to/file", "/path/to/file"},
		{"with dot dot", "/path/../to/file", "/to/file"},
		{"with dot", "/path/./to/file", "/path/to/file"},
		{"multiple slashes", "/path//to///file", "/path/to/file"},
		{"trailing slash", "/path/to/", "/path/to"},
		{"empty", "", "."},
		{"dot", ".", "."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.Clean(tt.input)
			if result != tt.expected {
				t.Errorf("Clean(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPath_Dir(t *testing.T) {
	p := NewPath()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple file", "/path/to/file.txt", "/path/to"},
		{"nested", "/a/b/c", "/a/b"},
		{"root file", "/file", "/"},
		{"just file", "file", "."},
		{"empty", "", "."},
		{"root", "/", "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.Dir(tt.input)
			if result != tt.expected {
				t.Errorf("Dir(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPath_Ext(t *testing.T) {
	p := NewPath()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"text file", "/path/to/file.txt", ".txt"},
		{"go file", "file.go", ".go"},
		{"no extension", "/path/to/file", ""},
		{"hidden file", ".bashrc", ".bashrc"},
		{"dot in path", "/path.txt/file", ""},
		{"multiple dots", "file.tar.gz", ".gz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.Ext(tt.input)
			if result != tt.expected {
				t.Errorf("Ext(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPath_IsAbs(t *testing.T) {
	p := NewPath()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"absolute", "/path/to/file", true},
		{"relative", "path/to/file", false},
		{"current dir", "./file", false},
		{"parent dir", "../file", false},
		{"root", "/", true},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.IsAbs(tt.input)
			if result != tt.expected {
				t.Errorf("IsAbs(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPath_Join(t *testing.T) {
	p := NewPath()

	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"two paths", []string{"/path", "to/file"}, "/path/to/file"},
		{"multiple paths", []string{"a", "b", "c", "d"}, "a/b/c/d"},
		{"with dots", []string{"/path", "..", "to", "file"}, "/to/file"},
		{"empty elements", []string{"/path", "", "file"}, "/path/file"},
		{"single element", []string{"/path"}, "/path"},
		{"empty", []string{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := p.Join(tt.input...)
			if result != tt.expected {
				t.Errorf("Join(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPath_Match(t *testing.T) {
	p := NewPath()

	tests := []struct {
		name        string
		pattern     string
		path        string
		expectMatch bool
		expectError bool
	}{
		{"simple match", "*.txt", "file.txt", true, false},
		{"no match", "*.txt", "file.go", false, false},
		{"question mark", "file?.txt", "file1.txt", true, false},
		{"range match", "file[0-9].txt", "file5.txt", true, false},
		{"range no match", "file[0-9].txt", "filea.txt", false, false},
		{"complex pattern", "/path/*/file.txt", "/path/to/file.txt", true, false},
		{"invalid pattern", "[]a]", "a", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched, err := p.Match(tt.pattern, tt.path)
			if tt.expectError {
				if err == nil {
					t.Errorf("Match(%q, %q) expected error, got nil", tt.pattern, tt.path)
				}
				return
			}
			if err != nil {
				t.Errorf("Match(%q, %q) unexpected error: %v", tt.pattern, tt.path, err)
				return
			}
			if matched != tt.expectMatch {
				t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.path, matched, tt.expectMatch)
			}
		})
	}
}

func TestPath_Split(t *testing.T) {
	p := NewPath()

	tests := []struct {
		name     string
		input    string
		wantDir  string
		wantFile string
	}{
		{"simple file", "/path/to/file.txt", "/path/to/", "file.txt"},
		{"root file", "/file", "/", "file"},
		{"just file", "file", "", "file"},
		{"trailing slash", "/path/to/", "/path/to/", ""},
		{"empty", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, file := p.Split(tt.input)
			if dir != tt.wantDir || file != tt.wantFile {
				t.Errorf("Split(%q) = (%q, %q), want (%q, %q)",
					tt.input, dir, file, tt.wantDir, tt.wantFile)
			}
		})
	}
}
