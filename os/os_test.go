package os

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNewOS(t *testing.T) {
	os := NewOS()
	_ = os
}

// File Operations

func TestOS_Create(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	defer file.Close()

	if file == nil {
		t.Fatal("Create() returned nil file")
	}
}

func TestOS_Open(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	// Create a file first
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	// Open it
	file, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	defer file.Close()

	if file == nil {
		t.Fatal("Open() returned nil file")
	}
}

func TestOS_ReadFileWriteFile(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	content := []byte("hello world")

	// Write file
	err := os.WriteFile(testFile, content, 0644)
	if err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	// Read file
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if string(data) != string(content) {
		t.Errorf("ReadFile() = %q, want %q", string(data), string(content))
	}
}

func TestOS_Remove(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	// Remove it
	err := os.Remove(testFile)
	if err != nil {
		t.Errorf("Remove() error = %v", err)
	}

	// Verify it's gone
	_, err = os.Stat(testFile)
	if err == nil {
		t.Error("File still exists after Remove()")
	}
}

func TestOS_Rename(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	oldPath := filepath.Join(tmpDir, "old.txt")
	newPath := filepath.Join(tmpDir, "new.txt")

	os.WriteFile(oldPath, []byte("test"), 0644)

	err := os.Rename(oldPath, newPath)
	if err != nil {
		t.Errorf("Rename() error = %v", err)
	}

	// Verify new file exists
	_, err = os.Stat(newPath)
	if err != nil {
		t.Errorf("Renamed file doesn't exist: %v", err)
	}
}

func TestOS_Truncate(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("hello world"), 0644)

	// Truncate to 5 bytes
	err := os.Truncate(testFile, 5)
	if err != nil {
		t.Errorf("Truncate() error = %v", err)
	}

	// Verify size
	info, _ := os.Stat(testFile)
	if info.Size() != 5 {
		t.Errorf("File size = %d, want 5", info.Size())
	}
}

func TestOS_Chmod(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	// Change mode
	err := os.Chmod(testFile, 0600)
	if err != nil {
		t.Errorf("Chmod() error = %v", err)
	}

	// Verify (may be platform-specific)
	info, _ := os.Stat(testFile)
	if info.Mode().Perm() != 0600 {
		t.Logf("Mode = %o, want 0600 (may vary by platform)", info.Mode().Perm())
	}
}

// Directory Operations

func TestOS_Mkdir(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testDir := filepath.Join(tmpDir, "testdir")
	err := os.Mkdir(testDir, 0755)
	if err != nil {
		t.Errorf("Mkdir() error = %v", err)
	}

	// Verify it exists
	info, err := os.Stat(testDir)
	if err != nil {
		t.Errorf("Stat() error = %v", err)
	}
	if !info.IsDir() {
		t.Error("Created path is not a directory")
	}
}

func TestOS_MkdirAll(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testDir := filepath.Join(tmpDir, "a", "b", "c")
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Errorf("MkdirAll() error = %v", err)
	}

	// Verify it exists
	info, err := os.Stat(testDir)
	if err != nil {
		t.Errorf("Stat() error = %v", err)
	}
	if !info.IsDir() {
		t.Error("Created path is not a directory")
	}
}

func TestOS_ReadDir(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	// Create some files
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "file2.txt"), []byte("test"), 0644)
	os.Mkdir(filepath.Join(tmpDir, "subdir"), 0755)

	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Errorf("ReadDir() error = %v", err)
		return
	}

	if len(entries) != 3 {
		t.Errorf("ReadDir() returned %d entries, want 3", len(entries))
	}
}

func TestOS_RemoveAll(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	// Create a directory tree
	testDir := filepath.Join(tmpDir, "remove_test")
	os.MkdirAll(filepath.Join(testDir, "a", "b"), 0755)
	os.WriteFile(filepath.Join(testDir, "file.txt"), []byte("test"), 0644)

	// Remove all
	err := os.RemoveAll(testDir)
	if err != nil {
		t.Errorf("RemoveAll() error = %v", err)
	}

	// Verify it's gone
	_, err = os.Stat(testDir)
	if err == nil {
		t.Error("Directory still exists after RemoveAll()")
	}
}

func TestOS_Chdir_Getwd(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	// Get current directory
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd() error = %v", err)
	}
	defer os.Chdir(oldWd) // Restore

	// Change to temp dir
	err = os.Chdir(tmpDir)
	if err != nil {
		t.Errorf("Chdir() error = %v", err)
	}

	// Verify
	newWd, err := os.Getwd()
	if err != nil {
		t.Errorf("Getwd() error = %v", err)
	}

	if newWd != tmpDir {
		t.Errorf("Getwd() = %q, want %q", newWd, tmpDir)
	}
}

// Environment Variables

func TestOS_Getenv_Setenv(t *testing.T) {
	os := NewOS()

	key := "TEST_ENV_VAR_12345"
	value := "test_value"

	// Set environment variable
	err := os.Setenv(key, value)
	if err != nil {
		t.Errorf("Setenv() error = %v", err)
	}
	defer os.Unsetenv(key)

	// Get it
	result := os.Getenv(key)
	if result != value {
		t.Errorf("Getenv() = %q, want %q", result, value)
	}
}

func TestOS_LookupEnv(t *testing.T) {
	os := NewOS()

	key := "TEST_ENV_VAR_67890"
	value := "test_value"

	// Set environment variable
	os.Setenv(key, value)
	defer os.Unsetenv(key)

	// Lookup
	result, ok := os.LookupEnv(key)
	if !ok {
		t.Error("LookupEnv() ok = false, want true")
	}
	if result != value {
		t.Errorf("LookupEnv() = %q, want %q", result, value)
	}

	// Lookup non-existent
	_, ok = os.LookupEnv("NONEXISTENT_VAR_12345")
	if ok {
		t.Error("LookupEnv() for non-existent var ok = true, want false")
	}
}

func TestOS_Unsetenv(t *testing.T) {
	os := NewOS()

	key := "TEST_ENV_VAR_UNSET"
	os.Setenv(key, "value")

	// Unset it
	err := os.Unsetenv(key)
	if err != nil {
		t.Errorf("Unsetenv() error = %v", err)
	}

	// Verify it's gone
	_, ok := os.LookupEnv(key)
	if ok {
		t.Error("Environment variable still exists after Unsetenv()")
	}
}

func TestOS_Environ(t *testing.T) {
	os := NewOS()

	environ := os.Environ()
	if len(environ) == 0 {
		t.Error("Environ() returned empty slice")
	}

	// Each entry should be in KEY=VALUE format
	for _, env := range environ {
		if !strings.Contains(env, "=") {
			t.Errorf("Environ() entry %q doesn't contain '='", env)
			break
		}
	}
}

func TestOS_ExpandEnv(t *testing.T) {
	os := NewOS()

	key := "TEST_EXPAND_VAR"
	value := "expanded_value"
	os.Setenv(key, value)
	defer os.Unsetenv(key)

	input := "${TEST_EXPAND_VAR}/path"
	result := os.ExpandEnv(input)

	expected := value + "/path"
	if result != expected {
		t.Errorf("ExpandEnv() = %q, want %q", result, expected)
	}
}

// File Info

func TestOS_Stat(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test content"), 0644)

	info, err := os.Stat(testFile)
	if err != nil {
		t.Errorf("Stat() error = %v", err)
		return
	}

	if info == nil {
		t.Fatal("Stat() returned nil FileInfo")
	}

	if info.Name() != "test.txt" {
		t.Errorf("Name() = %q, want %q", info.Name(), "test.txt")
	}

	if info.Size() != 12 {
		t.Errorf("Size() = %d, want 12", info.Size())
	}

	if info.IsDir() {
		t.Error("IsDir() = true for file, want false")
	}
}

func TestOS_Lstat(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	info, err := os.Lstat(testFile)
	if err != nil {
		t.Errorf("Lstat() error = %v", err)
		return
	}

	if info == nil {
		t.Fatal("Lstat() returned nil FileInfo")
	}
}

// Process Info

func TestOS_Getpid(t *testing.T) {
	os := NewOS()

	pid := os.Getpid()
	if pid <= 0 {
		t.Errorf("Getpid() = %d, want > 0", pid)
	}
}

func TestOS_Getppid(t *testing.T) {
	os := NewOS()

	ppid := os.Getppid()
	if ppid <= 0 {
		t.Errorf("Getppid() = %d, want > 0", ppid)
	}
}

func TestOS_Getuid(t *testing.T) {
	os := NewOS()

	uid := os.Getuid()
	// UID can be 0 (root), so just check it's not negative
	if uid < 0 {
		t.Errorf("Getuid() = %d, want >= 0", uid)
	}
}

func TestOS_Getgid(t *testing.T) {
	os := NewOS()

	gid := os.Getgid()
	if gid < 0 {
		t.Errorf("Getgid() = %d, want >= 0", gid)
	}
}

func TestOS_Geteuid(t *testing.T) {
	os := NewOS()

	euid := os.Geteuid()
	if euid < 0 {
		t.Errorf("Geteuid() = %d, want >= 0", euid)
	}
}

func TestOS_Getegid(t *testing.T) {
	os := NewOS()

	egid := os.Getegid()
	if egid < 0 {
		t.Errorf("Getegid() = %d, want >= 0", egid)
	}
}

func TestOS_Getpagesize(t *testing.T) {
	os := NewOS()

	pagesize := os.Getpagesize()
	if pagesize <= 0 {
		t.Errorf("Getpagesize() = %d, want > 0", pagesize)
	}
}

// Temporary Files/Directories

func TestOS_TempDir(t *testing.T) {
	os := NewOS()

	tmpDir := os.TempDir()
	if tmpDir == "" {
		t.Error("TempDir() returned empty string")
	}

	// Verify it exists
	info, err := os.Stat(tmpDir)
	if err != nil {
		t.Errorf("TempDir() returned non-existent directory: %v", err)
	}
	if info != nil && !info.IsDir() {
		t.Error("TempDir() returned non-directory")
	}
}

func TestOS_MkdirTemp(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	dir, err := os.MkdirTemp(tmpDir, "test-*")
	if err != nil {
		t.Errorf("MkdirTemp() error = %v", err)
		return
	}
	defer os.RemoveAll(dir)

	if dir == "" {
		t.Error("MkdirTemp() returned empty string")
	}

	// Verify it exists
	info, err := os.Stat(dir)
	if err != nil {
		t.Errorf("Stat() error = %v", err)
	}
	if !info.IsDir() {
		t.Error("MkdirTemp() didn't create a directory")
	}
}

func TestOS_CreateTemp(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	file, err := os.CreateTemp(tmpDir, "test-*.txt")
	if err != nil {
		t.Errorf("CreateTemp() error = %v", err)
		return
	}
	defer file.Close()
	defer os.Remove(file.Name())

	if file == nil {
		t.Fatal("CreateTemp() returned nil file")
	}

	name := file.Name()
	if name == "" {
		t.Error("CreateTemp() file has empty name")
	}
}

// Standard File Descriptors

func TestOS_Stdin_Stdout_Stderr(t *testing.T) {
	os := NewOS()

	stdin := os.Stdin()
	if stdin == nil {
		t.Error("Stdin() returned nil")
	}

	stdout := os.Stdout()
	if stdout == nil {
		t.Error("Stdout() returned nil")
	}

	stderr := os.Stderr()
	if stderr == nil {
		t.Error("Stderr() returned nil")
	}
}

// Miscellaneous

func TestOS_Hostname(t *testing.T) {
	os := NewOS()

	hostname, err := os.Hostname()
	if err != nil {
		t.Errorf("Hostname() error = %v", err)
		return
	}

	if hostname == "" {
		t.Error("Hostname() returned empty string")
	}
}

func TestOS_Executable(t *testing.T) {
	os := NewOS()

	exe, err := os.Executable()
	if err != nil {
		t.Errorf("Executable() error = %v", err)
		return
	}

	if exe == "" {
		t.Error("Executable() returned empty string")
	}
}

func TestOS_UserHomeDir(t *testing.T) {
	os := NewOS()

	dir, err := os.UserHomeDir()
	if err != nil {
		t.Errorf("UserHomeDir() error = %v", err)
		return
	}

	if dir == "" {
		t.Error("UserHomeDir() returned empty string")
	}
}

func TestOS_UserCacheDir(t *testing.T) {
	os := NewOS()

	dir, err := os.UserCacheDir()
	if err != nil {
		t.Errorf("UserCacheDir() error = %v", err)
		return
	}

	if dir == "" {
		t.Error("UserCacheDir() returned empty string")
	}
}

func TestOS_UserConfigDir(t *testing.T) {
	os := NewOS()

	dir, err := os.UserConfigDir()
	if err != nil {
		t.Errorf("UserConfigDir() error = %v", err)
		return
	}

	if dir == "" {
		t.Error("UserConfigDir() returned empty string")
	}
}

func TestOS_IsPathSeparator(t *testing.T) {
	os := NewOS()

	if !os.IsPathSeparator('/') {
		t.Error("IsPathSeparator('/') = false, want true")
	}

	if os.IsPathSeparator('a') {
		t.Error("IsPathSeparator('a') = true, want false")
	}
}

func TestOS_Pipe(t *testing.T) {
	os := NewOS()

	r, w, err := os.Pipe()
	if err != nil {
		t.Errorf("Pipe() error = %v", err)
		return
	}
	defer r.Close()
	defer w.Close()

	if r == nil || w == nil {
		t.Fatal("Pipe() returned nil file")
	}

	// Test writing and reading
	message := []byte("test pipe")
	go func() {
		w.Write(message)
		w.Close()
	}()

	buf := make([]byte, 100)
	n, err := r.Read(buf)
	if err != nil {
		t.Errorf("Read() error = %v", err)
		return
	}

	if string(buf[:n]) != string(message) {
		t.Errorf("Pipe read %q, want %q", string(buf[:n]), string(message))
	}
}

func TestOS_DirFS(t *testing.T) {
	os := NewOS()
	tmpDir := t.TempDir()

	// Create a test file
	os.WriteFile(filepath.Join(tmpDir, "test.txt"), []byte("test"), 0644)

	fsys := os.DirFS(tmpDir)
	if fsys == nil {
		t.Fatal("DirFS() returned nil")
	}

	// Try to open the file through the FS
	file, err := fsys.Open("test.txt")
	if err != nil {
		t.Errorf("FS.Open() error = %v", err)
		return
	}
	defer file.Close()

	if file == nil {
		t.Error("FS.Open() returned nil")
	}
}

func TestOS_Args(t *testing.T) {
	os := NewOS()

	args := os.Args()
	if len(args) == 0 {
		t.Error("Args() returned empty slice")
	}

	// First arg should be the test binary
	if args[0] == "" {
		t.Error("Args()[0] is empty")
	}
}
