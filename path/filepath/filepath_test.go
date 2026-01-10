package filepath

import (
	"os"
	"runtime"
	"testing"
)

func TestNewFilePath(t *testing.T) {
	fp := NewFilePath()
	_ = fp
}

func TestFilePath_Base(t *testing.T) {
	fp := NewFilePath()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple file", "/path/to/file.txt", "file.txt"},
		{"directory", "/path/to/dir/", "dir"},
		{"just filename", "file.txt", "file.txt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fp.Base(tt.input)
			if result != tt.expected {
				t.Errorf("Base(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFilePath_Clean(t *testing.T) {
	fp := NewFilePath()

	result := fp.Clean("/path//to/../file")
	// Just verify Clean returns something reasonable
	if result == "" {
		t.Error("Clean returned empty string")
	}
}

func TestFilePath_Dir(t *testing.T) {
	fp := NewFilePath()

	result := fp.Dir("/path/to/file.txt")
	if result == "" {
		t.Error("Dir returned empty string")
	}
}

func TestFilePath_Ext(t *testing.T) {
	fp := NewFilePath()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"text file", "file.txt", ".txt"},
		{"go file", "file.go", ".go"},
		{"no extension", "file", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fp.Ext(tt.input)
			if result != tt.expected {
				t.Errorf("Ext(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFilePath_IsAbs(t *testing.T) {
	fp := NewFilePath()

	// Test with platform-appropriate paths
	if runtime.GOOS == "windows" {
		if !fp.IsAbs("C:\\path\\to\\file") {
			t.Error("Expected Windows absolute path to be absolute")
		}
		if fp.IsAbs("path\\to\\file") {
			t.Error("Expected Windows relative path to be relative")
		}
	} else {
		if !fp.IsAbs("/path/to/file") {
			t.Error("Expected Unix absolute path to be absolute")
		}
		if fp.IsAbs("path/to/file") {
			t.Error("Expected Unix relative path to be relative")
		}
	}
}

func TestFilePath_Join(t *testing.T) {
	fp := NewFilePath()

	result := fp.Join("path", "to", "file")
	if result == "" {
		t.Error("Join returned empty string")
	}
}

func TestFilePath_Match(t *testing.T) {
	fp := NewFilePath()

	tests := []struct {
		name        string
		pattern     string
		path        string
		expectMatch bool
		expectError bool
	}{
		{"simple match", "*.txt", "file.txt", true, false},
		{"no match", "*.txt", "file.go", false, false},
		{"invalid pattern", "[]a]", "a", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matched, err := fp.Match(tt.pattern, tt.path)
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

func TestFilePath_Split(t *testing.T) {
	fp := NewFilePath()

	dir, file := fp.Split("/path/to/file.txt")
	if dir == "" && file == "" {
		t.Error("Split returned empty strings")
	}
}

func TestFilePath_Abs(t *testing.T) {
	fp := NewFilePath()

	// Test with current directory
	abs, err := fp.Abs(".")
	if err != nil {
		t.Errorf("Abs(\".\") returned error: %v", err)
	}
	if abs == "" {
		t.Error("Abs returned empty string")
	}
	if !fp.IsAbs(abs) {
		t.Errorf("Abs(\".\") = %q, expected absolute path", abs)
	}
}

func TestFilePath_FromSlash(t *testing.T) {
	fp := NewFilePath()

	result := fp.FromSlash("path/to/file")
	if result == "" {
		t.Error("FromSlash returned empty string")
	}
}

func TestFilePath_ToSlash(t *testing.T) {
	fp := NewFilePath()

	// ToSlash should always use forward slashes
	result := fp.ToSlash("path/to/file")
	if result != "path/to/file" {
		t.Errorf("ToSlash(\"path/to/file\") = %q, want \"path/to/file\"", result)
	}
}

func TestFilePath_Glob(t *testing.T) {
	fp := NewFilePath()

	// Create temp files for globbing
	tmpDir := t.TempDir()
	testFile1 := fp.Join(tmpDir, "file1.txt")
	testFile2 := fp.Join(tmpDir, "file2.txt")
	testFile3 := fp.Join(tmpDir, "file.go")

	// Create test files
	for _, file := range []string{testFile1, testFile2, testFile3} {
		if err := os.WriteFile(file, []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	// Test glob pattern
	pattern := fp.Join(tmpDir, "*.txt")
	matches, err := fp.Glob(pattern)
	if err != nil {
		t.Errorf("Glob(%q) returned error: %v", pattern, err)
	}
	if len(matches) != 2 {
		t.Errorf("Glob(%q) matched %d files, want 2", pattern, len(matches))
	}
}

func TestFilePath_IsLocal(t *testing.T) {
	fp := NewFilePath()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"simple path", "path/to/file", true},
		{"current dir", ".", true},
		{"parent dir", "..", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fp.IsLocal(tt.input)
			if result != tt.expected {
				t.Errorf("IsLocal(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFilePath_Localize(t *testing.T) {
	fp := NewFilePath()

	// Test with simple path
	result, err := fp.Localize("path/to/file")
	if err != nil {
		t.Errorf("Localize returned error: %v", err)
	}
	if result == "" {
		t.Error("Localize returned empty string")
	}
}

func TestFilePath_Rel(t *testing.T) {
	fp := NewFilePath()

	// Create temp directory for testing
	tmpDir := t.TempDir()
	subDir := fp.Join(tmpDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	rel, err := fp.Rel(tmpDir, subDir)
	if err != nil {
		t.Errorf("Rel(%q, %q) returned error: %v", tmpDir, subDir, err)
	}
	if rel != "sub" {
		t.Errorf("Rel(%q, %q) = %q, want \"sub\"", tmpDir, subDir, rel)
	}
}

func TestFilePath_SplitList(t *testing.T) {
	fp := NewFilePath()

	// Platform-specific path separator
	var pathList string
	if runtime.GOOS == "windows" {
		pathList = "C:\\path1;C:\\path2"
	} else {
		pathList = "/path1:/path2"
	}

	result := fp.SplitList(pathList)
	if len(result) != 2 {
		t.Errorf("SplitList(%q) returned %d elements, want 2", pathList, len(result))
	}
}

func TestFilePath_VolumeName(t *testing.T) {
	fp := NewFilePath()

	if runtime.GOOS == "windows" {
		vol := fp.VolumeName("C:\\path\\to\\file")
		if vol != "C:" {
			t.Errorf("VolumeName(\"C:\\\\path\\\\to\\\\file\") = %q, want \"C:\"", vol)
		}
	} else {
		vol := fp.VolumeName("/path/to/file")
		if vol != "" {
			t.Errorf("VolumeName(\"/path/to/file\") = %q, want \"\"", vol)
		}
	}
}

func TestFilePath_Walk(t *testing.T) {
	fp := NewFilePath()

	// Create temp directory structure
	tmpDir := t.TempDir()
	subDir := fp.Join(tmpDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	testFile := fp.Join(subDir, "file.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Walk the directory
	count := 0
	err := fp.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		count++
		return nil
	})

	if err != nil {
		t.Errorf("Walk returned error: %v", err)
	}
	if count < 3 { // tmpDir, subDir, testFile
		t.Errorf("Walk visited %d paths, want at least 3", count)
	}
}

func TestFilePath_WalkDir(t *testing.T) {
	fp := NewFilePath()

	// Create temp directory structure
	tmpDir := t.TempDir()
	subDir := fp.Join(tmpDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	testFile := fp.Join(subDir, "file.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Walk the directory
	count := 0
	err := fp.WalkDir(tmpDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		count++
		return nil
	})

	if err != nil {
		t.Errorf("WalkDir returned error: %v", err)
	}
	if count < 3 { // tmpDir, subDir, testFile
		t.Errorf("WalkDir visited %d paths, want at least 3", count)
	}
}

func TestFilePath_EvalSymlinks(t *testing.T) {
	fp := NewFilePath()

	// Test with current directory
	result, err := fp.EvalSymlinks(".")
	if err != nil {
		t.Errorf("EvalSymlinks(\".\") returned error: %v", err)
	}
	if result == "" {
		t.Error("EvalSymlinks returned empty string")
	}
}
