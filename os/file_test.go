package os

import (
	"io"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestFile_ReadWrite(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	// Write
	message := []byte("hello world")
	n, err := file.Write(message)
	if err != nil {
		t.Errorf("Write() error = %v", err)
	}
	if n != len(message) {
		t.Errorf("Write() wrote %d bytes, want %d", n, len(message))
	}

	// Seek to beginning
	_, err = file.Seek(0, 0)
	if err != nil {
		t.Errorf("Seek() error = %v", err)
	}

	// Read
	buf := make([]byte, 100)
	n, err = file.Read(buf)
	if err != nil && err != io.EOF {
		t.Errorf("Read() error = %v", err)
	}

	if string(buf[:n]) != string(message) {
		t.Errorf("Read() = %q, want %q", string(buf[:n]), string(message))
	}
}

func TestFile_WriteString(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	message := "hello string"
	n, err := file.WriteString(message)
	if err != nil {
		t.Errorf("WriteString() error = %v", err)
	}
	if n != len(message) {
		t.Errorf("WriteString() wrote %d bytes, want %d", n, len(message))
	}
}

func TestFile_ReadAtWriteAt(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	// WriteAt
	message := []byte("test")
	n, err := file.WriteAt(message, 5)
	if err != nil {
		t.Errorf("WriteAt() error = %v", err)
	}
	if n != len(message) {
		t.Errorf("WriteAt() wrote %d bytes, want %d", n, len(message))
	}

	// ReadAt
	buf := make([]byte, 4)
	n, err = file.ReadAt(buf, 5)
	if err != nil && err != io.EOF {
		t.Errorf("ReadAt() error = %v", err)
	}
	if string(buf) != "test" {
		t.Errorf("ReadAt() = %q, want %q", string(buf), "test")
	}
}

func TestFile_Seek(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("hello world"), 0644)

	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer file.Close()

	// Seek to position 6
	pos, err := file.Seek(6, 0)
	if err != nil {
		t.Errorf("Seek() error = %v", err)
	}
	if pos != 6 {
		t.Errorf("Seek() = %d, want 6", pos)
	}

	// Read from there
	buf := make([]byte, 5)
	n, _ := file.Read(buf)
	if string(buf[:n]) != "world" {
		t.Errorf("Read after Seek() = %q, want %q", string(buf[:n]), "world")
	}
}

func TestFile_Name(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	name := file.Name()
	if name != testFile {
		t.Errorf("Name() = %q, want %q", name, testFile)
	}
}

func TestFile_Fd(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	fd := file.Fd()
	if fd == 0 {
		t.Error("Fd() = 0, want non-zero")
	}
}

func TestFile_Stat(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		t.Errorf("Stat() error = %v", err)
		return
	}

	if info.Name() != "test.txt" {
		t.Errorf("Stat().Name() = %q, want %q", info.Name(), "test.txt")
	}

	if info.Size() != 12 {
		t.Errorf("Stat().Size() = %d, want 12", info.Size())
	}
}

func TestFile_Truncate(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	// Write some data
	file.WriteString("hello world")

	// Truncate to 5 bytes
	err = file.Truncate(5)
	if err != nil {
		t.Errorf("Truncate() error = %v", err)
	}

	// Check size
	info, _ := file.Stat()
	if info.Size() != 5 {
		t.Errorf("Size after Truncate() = %d, want 5", info.Size())
	}
}

func TestFile_Sync(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	file.WriteString("test")

	err = file.Sync()
	if err != nil {
		t.Errorf("Sync() error = %v", err)
	}
}

func TestFile_Close(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	err = file.Close()
	if err != nil {
		t.Errorf("Close() error = %v", err)
	}

	// Writing after close should fail
	_, err = file.Write([]byte("test"))
	if err == nil {
		t.Error("Write() after Close() should fail")
	}
}

func TestFile_Chmod(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	err = file.Chmod(0600)
	if err != nil {
		t.Errorf("Chmod() error = %v", err)
	}

	// Verify (may be platform-specific)
	info, _ := file.Stat()
	if info.Mode().Perm() != 0600 {
		t.Logf("Mode = %o, want 0600 (may vary by platform)", info.Mode().Perm())
	}
}

func TestFile_ReadDir(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	// Create some files in the directory
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "file2.txt"), []byte("test"), 0644)

	// Open the directory
	dir, err := os.Open(tmpDir)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer dir.Close()

	// ReadDir
	entries, err := dir.ReadDir(-1)
	if err != nil {
		t.Errorf("ReadDir() error = %v", err)
		return
	}

	if len(entries) < 2 {
		t.Errorf("ReadDir() returned %d entries, want at least 2", len(entries))
	}
}

func TestFile_Readdir(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	// Create some files in the directory
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "file2.txt"), []byte("test"), 0644)

	// Open the directory
	dir, err := os.Open(tmpDir)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer dir.Close()

	// Readdir
	infos, err := dir.Readdir(-1)
	if err != nil {
		t.Errorf("Readdir() error = %v", err)
		return
	}

	if len(infos) < 2 {
		t.Errorf("Readdir() returned %d infos, want at least 2", len(infos))
	}
}

func TestFile_Readdirnames(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	// Create some files in the directory
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "file2.txt"), []byte("test"), 0644)

	// Open the directory
	dir, err := os.Open(tmpDir)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer dir.Close()

	// Readdirnames
	names, err := dir.Readdirnames(-1)
	if err != nil {
		t.Errorf("Readdirnames() error = %v", err)
		return
	}

	if len(names) < 2 {
		t.Errorf("Readdirnames() returned %d names, want at least 2", len(names))
	}
}

func TestFile_ReadFrom(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	// ReadFrom a strings.Reader
	reader := strings.NewReader("hello from reader")
	n, err := file.ReadFrom(reader)
	if err != nil {
		t.Errorf("ReadFrom() error = %v", err)
	}

	if n != 17 {
		t.Errorf("ReadFrom() copied %d bytes, want 17", n)
	}

	// Verify contents
	file.Seek(0, 0)
	buf := make([]byte, 100)
	nread, _ := file.Read(buf)
	if string(buf[:nread]) != "hello from reader" {
		t.Errorf("File contents = %q, want %q", string(buf[:nread]), "hello from reader")
	}
}

func TestFile_WriteTo(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer file.Close()

	// WriteTo a strings.Builder
	var builder strings.Builder
	n, err := file.WriteTo(&builder)
	if err != nil {
		t.Errorf("WriteTo() error = %v", err)
	}

	if n != 12 {
		t.Errorf("WriteTo() wrote %d bytes, want 12", n)
	}

	if builder.String() != "test content" {
		t.Errorf("WriteTo() result = %q, want %q", builder.String(), "test content")
	}
}

func TestFile_Deadlines(t *testing.T) {
	os := NewOS()

	// Create a pipe for testing deadlines
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Pipe() error = %v", err)
	}
	defer r.Close()
	defer w.Close()

	// Set deadlines (may not work on all file types)
	deadline := time.Now().Add(1 * time.Second)

	err = r.SetReadDeadline(deadline)
	if err != nil {
		// Some file types don't support deadlines
		t.Logf("SetReadDeadline() error (may be expected): %v", err)
	}

	err = w.SetWriteDeadline(deadline)
	if err != nil {
		t.Logf("SetWriteDeadline() error (may be expected): %v", err)
	}

	err = r.SetDeadline(deadline)
	if err != nil {
		t.Logf("SetDeadline() error (may be expected): %v", err)
	}
}

func TestFile_SyscallConn(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	conn, err := file.SyscallConn()
	if err != nil {
		t.Errorf("SyscallConn() error = %v", err)
		return
	}

	if conn == nil {
		t.Error("SyscallConn() returned nil")
	}
}

func TestFile_Chdir(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	// Save current directory
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd() error = %v", err)
	}
	defer os.Chdir(oldWd)

	// Open the temp directory
	dir, err := os.Open(tmpDir)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer dir.Close()

	// Chdir to it
	err = dir.Chdir()
	if err != nil {
		t.Errorf("Chdir() error = %v", err)
	}

	// Verify
	newWd, err := os.Getwd()
	if err != nil {
		t.Errorf("Getwd() error = %v", err)
	}

	if newWd != tmpDir {
		t.Errorf("Getwd() after Chdir() = %q, want %q", newWd, tmpDir)
	}
}
