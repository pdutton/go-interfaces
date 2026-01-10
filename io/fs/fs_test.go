package fs

import (
	stdfs "io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestNewFileSystem(t *testing.T) {
	fs := NewFileSystem()
	_ = fs
}

func TestFileSystem_ValidPath(t *testing.T) {
	fs := NewFileSystem()

	tests := []struct {
		name  string
		path  string
		valid bool
	}{
		{"valid path", "test/path.txt", true},
		{"valid simple", "file.txt", true},
		{"invalid absolute", "/absolute/path", false},
		{"invalid dot dot", "../parent", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fs.ValidPath(tt.path)
			if result != tt.valid {
				t.Errorf("ValidPath(%q) = %v, want %v", tt.path, result, tt.valid)
			}
		})
	}
}

func TestFileSystem_ReadFile(t *testing.T) {
	fs := NewFileSystem()

	// Create a temp directory with test files
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	testContent := []byte("hello world")
	if err := os.WriteFile(testFile, testContent, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create a DirFS
	dirFS := os.DirFS(tmpDir)

	// Read the file
	data, err := fs.ReadFile(dirFS, "test.txt")
	if err != nil {
		t.Errorf("ReadFile() error = %v", err)
		return
	}

	if string(data) != string(testContent) {
		t.Errorf("ReadFile() = %q, want %q", string(data), string(testContent))
	}
}

func TestFileSystem_ReadDir(t *testing.T) {
	fs := NewFileSystem()

	// Create a temp directory with test files
	tmpDir := t.TempDir()
	for i, name := range []string{"file1.txt", "file2.txt", "dir1"} {
		if i < 2 {
			if err := os.WriteFile(filepath.Join(tmpDir, name), []byte("test"), 0644); err != nil {
				t.Fatalf("Failed to create test file: %v", err)
			}
		} else {
			if err := os.Mkdir(filepath.Join(tmpDir, name), 0755); err != nil {
				t.Fatalf("Failed to create test directory: %v", err)
			}
		}
	}

	// Create a DirFS
	dirFS := os.DirFS(tmpDir)

	// Read the directory
	entries, err := fs.ReadDir(dirFS, ".")
	if err != nil {
		t.Errorf("ReadDir() error = %v", err)
		return
	}

	if len(entries) != 3 {
		t.Errorf("ReadDir() returned %d entries, want 3", len(entries))
	}
}

func TestFileSystem_Glob(t *testing.T) {
	fs := NewFileSystem()

	// Create a temp directory with test files
	tmpDir := t.TempDir()
	for _, name := range []string{"test1.txt", "test2.txt", "other.log"} {
		if err := os.WriteFile(filepath.Join(tmpDir, name), []byte("test"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	// Create a DirFS
	dirFS := os.DirFS(tmpDir)

	// Glob for txt files
	matches, err := fs.Glob(dirFS, "*.txt")
	if err != nil {
		t.Errorf("Glob() error = %v", err)
		return
	}

	if len(matches) != 2 {
		t.Errorf("Glob() matched %d files, want 2", len(matches))
	}
}

func TestFileSystem_WalkDir(t *testing.T) {
	fsys := NewFileSystem()

	// Create a temp directory with nested structure
	tmpDir := t.TempDir()
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("test"), 0644)
	os.Mkdir(filepath.Join(tmpDir, "subdir"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "subdir", "file2.txt"), []byte("test"), 0644)

	// Create a DirFS
	dirFS := os.DirFS(tmpDir)

	// Walk the directory
	var count int
	err := fsys.WalkDir(dirFS, ".", func(path string, d stdfs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		count++
		return nil
	})

	if err != nil {
		t.Errorf("WalkDir() error = %v", err)
	}

	// Should visit: . (root), file1.txt, subdir, subdir/file2.txt
	if count < 3 {
		t.Errorf("WalkDir() visited %d entries, want at least 3", count)
	}
}

func TestFileSystem_Stat(t *testing.T) {
	fs := NewFileSystem()

	// Create a temp directory with a test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("hello"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create a DirFS
	dirFS := os.DirFS(tmpDir)

	// Stat the file
	info, err := fs.Stat(dirFS, "test.txt")
	if err != nil {
		t.Errorf("Stat() error = %v", err)
		return
	}

	if info.Name() != "test.txt" {
		t.Errorf("Stat() Name() = %q, want %q", info.Name(), "test.txt")
	}

	if info.Size() != 5 {
		t.Errorf("Stat() Size() = %d, want 5", info.Size())
	}

	if info.IsDir() {
		t.Error("Stat() IsDir() = true, want false")
	}
}

func TestFileSystem_Sub(t *testing.T) {
	fs := NewFileSystem()

	// Create a temp directory with subdirectory
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "subdir")
	os.Mkdir(subDir, 0755)
	os.WriteFile(filepath.Join(subDir, "test.txt"), []byte("test"), 0644)

	// Create a DirFS
	dirFS := os.DirFS(tmpDir)

	// Create a sub filesystem
	subFS, err := fs.Sub(dirFS, "subdir")
	if err != nil {
		t.Errorf("Sub() error = %v", err)
		return
	}

	if subFS == nil {
		t.Error("Sub() returned nil filesystem")
	}

	// Verify we can read from the sub filesystem
	data, err := fs.ReadFile(subFS, "test.txt")
	if err != nil {
		t.Errorf("ReadFile() on sub filesystem error = %v", err)
		return
	}

	if string(data) != "test" {
		t.Errorf("ReadFile() on sub filesystem = %q, want %q", string(data), "test")
	}
}

func TestDirEntry_Methods(t *testing.T) {
	// Create a temp directory with a test file and directory
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)
	testDir := filepath.Join(tmpDir, "testdir")
	os.Mkdir(testDir, 0755)

	dirFS := os.DirFS(tmpDir)
	fs := NewFileSystem()

	entries, err := fs.ReadDir(dirFS, ".")
	if err != nil {
		t.Fatalf("ReadDir() error = %v", err)
	}

	if len(entries) < 2 {
		t.Fatalf("ReadDir() returned %d entries, want at least 2", len(entries))
	}

	// Test file entry
	var fileEntry, dirEntry DirEntry
	for _, entry := range entries {
		if entry.Name() == "test.txt" {
			fileEntry = entry
		} else if entry.Name() == "testdir" {
			dirEntry = entry
		}
	}

	if fileEntry == nil {
		t.Fatal("File entry not found")
	}

	// Test Name
	if fileEntry.Name() != "test.txt" {
		t.Errorf("Name() = %q, want %q", fileEntry.Name(), "test.txt")
	}

	// Test IsDir
	if fileEntry.IsDir() {
		t.Error("IsDir() = true for file, want false")
	}

	// Test Type
	fmType := fileEntry.Type()
	if fmType.IsDir() {
		t.Error("Type().IsDir() = true for file, want false")
	}

	// Test Info
	info, err := fileEntry.Info()
	if err != nil {
		t.Errorf("Info() error = %v", err)
	}
	if info.Name() != "test.txt" {
		t.Errorf("Info().Name() = %q, want %q", info.Name(), "test.txt")
	}

	// Test Format
	format := fileEntry.Format()
	if format == "" {
		t.Error("Format() returned empty string")
	}

	// Test Nub
	nub := fileEntry.Nub()
	if nub == nil {
		t.Error("Nub() returned nil")
	}

	// Test directory entry
	if dirEntry != nil {
		if !dirEntry.IsDir() {
			t.Error("IsDir() = false for directory, want true")
		}
	}
}

func TestFileInfo_Methods(t *testing.T) {
	// Create a temp file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	testContent := []byte("hello world")
	os.WriteFile(testFile, testContent, 0644)

	dirFS := os.DirFS(tmpDir)
	fs := NewFileSystem()

	info, err := fs.Stat(dirFS, "test.txt")
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}

	// Test Name
	if info.Name() != "test.txt" {
		t.Errorf("Name() = %q, want %q", info.Name(), "test.txt")
	}

	// Test Size
	if info.Size() != int64(len(testContent)) {
		t.Errorf("Size() = %d, want %d", info.Size(), len(testContent))
	}

	// Test Mode
	mode := info.Mode()
	if mode.IsDir() {
		t.Error("Mode().IsDir() = true for file, want false")
	}

	// Test ModTime
	modTime := info.ModTime()
	if modTime.IsZero() {
		t.Error("ModTime() returned zero time")
	}

	// Test IsDir
	if info.IsDir() {
		t.Error("IsDir() = true for file, want false")
	}

	// Test Sys
	_ = info.Sys()

	// Test Nub
	nub := info.Nub()
	if nub == nil {
		t.Error("Nub() returned nil")
	}
}

func TestFileMode_Methods(t *testing.T) {
	// Test regular file mode
	regularMode := NewFileMode(0644)

	if !regularMode.IsRegular() {
		t.Error("IsRegular() = false for regular file, want true")
	}

	if regularMode.IsDir() {
		t.Error("IsDir() = true for regular file, want false")
	}

	// Test directory mode
	dirMode := ModeDir

	if !dirMode.IsDir() {
		t.Error("IsDir() = false for directory, want true")
	}

	if dirMode.IsRegular() {
		t.Error("IsRegular() = true for directory, want false")
	}

	// Test symlink mode
	symlinkMode := ModeSymlink

	if !symlinkMode.IsSymlink() {
		t.Error("ModeSymlink.IsSymlink() = false, want true")
	}

	// Test named pipe mode
	namedPipeMode := ModeNamedPipe

	if !namedPipeMode.IsNamedPipe() {
		t.Error("ModeNamedPipe.IsNamedPipe() = false, want true")
	}

	// Test device mode
	deviceMode := ModeDevice

	if !deviceMode.IsDevice() {
		t.Error("ModeDevice.IsDevice() = false, want true")
	}

	// Test socket mode
	socketMode := ModeSocket

	if !socketMode.IsSocket() {
		t.Error("ModeSocket.IsSocket() = false, want true")
	}

	// Test char device mode
	charDeviceMode := ModeCharDevice

	if !charDeviceMode.IsCharDevice() {
		t.Error("ModeCharDevice.IsCharDevice() = false, want true")
	}

	// Test irregular mode
	irregularMode := ModeIrregular

	if !irregularMode.IsIrregular() {
		t.Error("ModeIrregular.IsIrregular() = false, want true")
	}

	// Test Perm
	perm := regularMode.Perm()
	if perm != 0644 {
		t.Errorf("Perm() = %o, want %o", perm, 0644)
	}

	// Test String
	str := regularMode.String()
	if str == "" {
		t.Error("String() returned empty string")
	}

	// Test Nub
	nub := regularMode.Nub()
	if nub != 0644 {
		t.Errorf("Nub() = %o, want %o", nub, 0644)
	}
}

func TestFileSystem_FileInfoToDirEntry(t *testing.T) {
	fs := NewFileSystem()

	// Create a temp file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	dirFS := os.DirFS(tmpDir)

	// Get FileInfo
	info, err := fs.Stat(dirFS, "test.txt")
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}

	// Convert to DirEntry
	entry := fs.FileInfoToDirEntry(info)
	if entry == nil {
		t.Fatal("FileInfoToDirEntry() returned nil")
	}

	if entry.Name() != "test.txt" {
		t.Errorf("FileInfoToDirEntry() Name() = %q, want %q", entry.Name(), "test.txt")
	}

	if entry.IsDir() {
		t.Error("FileInfoToDirEntry() IsDir() = true for file, want false")
	}
}

func TestFileSystem_FormatDirEntry(t *testing.T) {
	fs := NewFileSystem()

	// Create a temp directory with a test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	dirFS := os.DirFS(tmpDir)

	entries, err := fs.ReadDir(dirFS, ".")
	if err != nil {
		t.Fatalf("ReadDir() error = %v", err)
	}

	if len(entries) == 0 {
		t.Fatal("ReadDir() returned no entries")
	}

	// Format the first entry
	formatted := fs.FormatDirEntry(entries[0])
	if formatted == "" {
		t.Error("FormatDirEntry() returned empty string")
	}
}

func TestFileSystem_FormatFileInfo(t *testing.T) {
	fs := NewFileSystem()

	// Create a temp file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	dirFS := os.DirFS(tmpDir)

	info, err := fs.Stat(dirFS, "test.txt")
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}

	// Format the file info
	formatted := fs.FormatFileInfo(info)
	if formatted == "" {
		t.Error("FormatFileInfo() returned empty string")
	}
}

func TestNewDirEntry(t *testing.T) {
	// Create a temp directory with a file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(testFile, []byte("test"), 0644)

	dirFS := os.DirFS(tmpDir)

	// Read the directory to get stdlib DirEntry
	stdEntries, err := stdfs.ReadDir(dirFS, ".")
	if err != nil {
		t.Fatalf("ReadDir() error = %v", err)
	}

	if len(stdEntries) == 0 {
		t.Fatal("No entries found")
	}

	// Wrap the stdlib DirEntry
	entry := NewDirEntry(stdEntries[0])

	if entry.Name() != "test.txt" {
		t.Errorf("Name() = %q, want %q", entry.Name(), "test.txt")
	}
}

func TestNewDirEntryList(t *testing.T) {
	// Create a temp directory with files
	tmpDir := t.TempDir()
	os.WriteFile(filepath.Join(tmpDir, "file1.txt"), []byte("test"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "file2.txt"), []byte("test"), 0644)

	dirFS := os.DirFS(tmpDir)

	// Read the directory to get stdlib DirEntry slice
	stdEntries, err := stdfs.ReadDir(dirFS, ".")
	if err != nil {
		t.Fatalf("ReadDir() error = %v", err)
	}

	// Convert to facade list
	entries := NewDirEntryList(stdEntries)

	if len(entries) != len(stdEntries) {
		t.Errorf("NewDirEntryList() returned %d entries, want %d", len(entries), len(stdEntries))
	}

	for i, entry := range entries {
		if entry.Name() != stdEntries[i].Name() {
			t.Errorf("Entry %d: Name() = %q, want %q", i, entry.Name(), stdEntries[i].Name())
		}
	}
}

func TestNewFileInfoList(t *testing.T) {
	// Create a temp directory with files
	tmpDir := t.TempDir()
	file1 := filepath.Join(tmpDir, "file1.txt")
	file2 := filepath.Join(tmpDir, "file2.txt")
	os.WriteFile(file1, []byte("test"), 0644)
	os.WriteFile(file2, []byte("test"), 0644)

	// Get stdlib FileInfo objects
	info1, err := os.Stat(file1)
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}
	info2, err := os.Stat(file2)
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}

	stdInfos := []stdfs.FileInfo{info1, info2}

	// Convert to facade list
	infos := NewFileInfoList(stdInfos)

	if len(infos) != len(stdInfos) {
		t.Errorf("NewFileInfoList() returned %d infos, want %d", len(infos), len(stdInfos))
	}

	for i, info := range infos {
		if info.Name() != stdInfos[i].Name() {
			t.Errorf("Info %d: Name() = %q, want %q", i, info.Name(), stdInfos[i].Name())
		}
	}
}

