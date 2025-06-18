package os

import (
	"os"
	"time"

	"github.com/pdutton/go-interfaces/io/fs"
)

type OS interface {
	// Access to global variables:
	Stdin() File
	Stderr() File
	Stdout() File

	Args() []string

	// Functions:
	Chdir(string) error
	Chmod(string, FileMode) error
	Chown(string, int, int) error
	Chtimes(string, time.Time, time.Time) error
	Clearenv()
	CopyFS(string, fs.FS) error
	DirFS(string) fs.FS
	Environ() []string
	Executable() (string, error)
	Exit(int)
	Expand(string, func(string) string) string
	ExpandEnv(string) string
	Getegid() int
	Getenv(string) string
	Geteuid() int
	Getgid() int
	Getgroups() ([]int, error)
	Getpagesize() int
	Getpid() int
	Getppid() int
	Getuid() int
	Getwd() (string, error)
	Hostname() (string, error)
	IsExist(error) bool    // Deprecated
	IsNotExist(error) bool // Deprecated
	IsPathSeparator(uint8) bool
	IsPermission(error) bool // Deprecated
	IsTimeout(error) bool    // Deprecated
	Lchown(string, int, int) error
	Link(string, string) error
	LookupEnv(string) (string, bool)
	Mkdir(string, FileMode) error
	MkdirAll(string, FileMode) error
	MkdirTemp(string, string) (string, error)
	NewSyscallError(string, error) error
	Pipe() (File, File, error)
	ReadFile(string) ([]byte, error)
	Readlink(string) (string, error)
	Remove(string) error
	RemoveAll(string) error
	Rename(string, string) error
	SameFile(FileInfo, FileInfo) bool
	Setenv(string, string) error
	Symlink(string, string) error
	TempDir() string
	Truncate(string, int64) error
	Unsetenv(string) error
	UserCacheDir() (string, error)
	UserConfigDir() (string, error)
	UserHomeDir() (string, error)
	WriteFile(string, []byte, FileMode) error

	// DirEntry constructors
	ReadDir(string) ([]DirEntry, error)

	// File constructors
	Create(string) (File, error)
	CreateTemp(string, string) (File, error)
	NewFile(uintptr, string) File
	Open(string) (File, error)
	OpenFile(string, int, FileMode) (File, error)
	OpenInRoot(string, string) (File, error)

	// FileInfo constructors
	Lstat(string) (FileInfo, error)
	Stat(string) (FileInfo, error)

	// Process constructors
	FindProcess(int) (Process, error)
	StartProcess(string, []string, *ProcAttr) (Process, error)

	// Root constructors
	OpenRoot(string) (Root, error)
}

type osFacade struct{}

func NewOS() OS {
	return osFacade{}
}

func (_ osFacade) Stdin() File {
	return WrapFile(os.Stdin)
}

func (_ osFacade) Stderr() File {
	return WrapFile(os.Stderr)
}

func (_ osFacade) Stdout() File {
	return WrapFile(os.Stdout)
}

func (_ osFacade) Args() []string {
	return os.Args
}

func (_ osFacade) Chdir(dir string) error {
	return os.Chdir(dir)
}

func (_ osFacade) Chmod(name string, mode FileMode) error {
	return os.Chmod(name, mode.Nub())
}

func (_ osFacade) Chown(name string, uid, gid int) error {
	return os.Chown(name, uid, gid)
}

func (_ osFacade) Chtimes(name string, atime, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}

func (_ osFacade) Clearenv() {
	os.Clearenv()
}

func (_ osFacade) CopyFS(dir string, fsys fs.FS) error {
	return os.CopyFS(dir, fsys)
}

func (_ osFacade) DirFS(dir string) fs.FS {
	return os.DirFS(dir)
}

func (_ osFacade) Environ() []string {
	return os.Environ()
}

func (_ osFacade) Executable() (string, error) {
	return os.Executable()
}

func (_ osFacade) Exit(code int) {
	os.Exit(code)
}

func (_ osFacade) Expand(s string, mapping func(string) string) string {
	return os.Expand(s, mapping)
}

func (_ osFacade) ExpandEnv(s string) string {
	return os.ExpandEnv(s)
}

func (_ osFacade) Getegid() int {
	return os.Getegid()
}

func (_ osFacade) Getenv(key string) string {
	return os.Getenv(key)
}

func (_ osFacade) Geteuid() int {
	return os.Geteuid()
}

func (_ osFacade) Getgid() int {
	return os.Getgid()
}

func (_ osFacade) Getgroups() ([]int, error) {
	return os.Getgroups()
}

func (_ osFacade) Getpagesize() int {
	return os.Getpagesize()
}

func (_ osFacade) Getpid() int {
	return os.Getpid()
}

func (_ osFacade) Getppid() int {
	return os.Getppid()
}

func (_ osFacade) Getuid() int {
	return os.Getuid()
}

func (_ osFacade) Getwd() (string, error) {
	return os.Getwd()
}

func (_ osFacade) Hostname() (string, error) {
	return os.Hostname()
}

func (_ osFacade) IsExist(err error) bool {
	return os.IsExist(err)
}

func (_ osFacade) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (_ osFacade) IsPathSeparator(c uint8) bool {
	return os.IsPathSeparator(c)
}

func (_ osFacade) IsPermission(err error) bool {
	return os.IsPermission(err)
}

func (_ osFacade) IsTimeout(err error) bool {
	return os.IsTimeout(err)
}

func (_ osFacade) Lchown(name string, uid, gid int) error {
	return os.Lchown(name, uid, gid)
}

func (_ osFacade) Link(oldname, newname string) error {
	return os.Link(oldname, newname)
}

func (_ osFacade) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (_ osFacade) Mkdir(name string, perm FileMode) error {
	return os.Mkdir(name, perm.Nub())
}

func (_ osFacade) MkdirAll(name string, perm FileMode) error {
	return os.MkdirAll(name, perm.Nub())
}

func (_ osFacade) MkdirTemp(dir, pattern string) (string, error) {
	return os.MkdirTemp(dir, pattern)
}

func (_ osFacade) NewSyscallError(syscall string, err error) error {
	return os.NewSyscallError(syscall, err)
}

func (_ osFacade) Pipe() (File, File, error) {
	f1, f2, err := os.Pipe()
	return WrapFile(f1), WrapFile(f2), err
}

func (_ osFacade) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (_ osFacade) Readlink(name string) (string, error) {
	return os.Readlink(name)
}

func (_ osFacade) Remove(name string) error {
	return os.Remove(name)
}

func (_ osFacade) RemoveAll(name string) error {
	return os.RemoveAll(name)
}

func (_ osFacade) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (_ osFacade) SameFile(fi1, fi2 FileInfo) bool {
	return os.SameFile(fi1.Nub(), fi2.Nub())
}

func (_ osFacade) Setenv(key, value string) error {
	return os.Setenv(key, value)
}

func (_ osFacade) Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}

func (_ osFacade) TempDir() string {
	return os.TempDir()
}

func (_ osFacade) Truncate(name string, size int64) error {
	return os.Truncate(name, size)
}

func (_ osFacade) Unsetenv(key string) error {
	return os.Unsetenv(key)
}

func (_ osFacade) UserCacheDir() (string, error) {
	return os.UserCacheDir()
}

func (_ osFacade) UserConfigDir() (string, error) {
	return os.UserConfigDir()
}

func (_ osFacade) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func (_ osFacade) WriteFile(name string, data []byte, perm FileMode) error {
	return os.WriteFile(name, data, perm.Nub())
}

func (_ osFacade) ReadDir(name string) ([]DirEntry, error) {
	dea, err := os.ReadDir(name)
	return fs.NewDirEntryList(dea), err
}

func (_ osFacade) Lstat(name string) (FileInfo, error) {
	fi, err := os.Lstat(name)
	return fs.NewFileInfo(fi), err
}

func (_ osFacade) Stat(name string) (FileInfo, error) {
	fi, err := os.Stat(name)
	return fs.NewFileInfo(fi), err
}
