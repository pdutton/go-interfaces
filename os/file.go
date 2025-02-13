package os

import (
	"io"
	"os"
	"time"
)

type File interface {
	Chdir() error
	Chmod(FileMode) error
	Chown(int, int) error
	Close() error
	Fd() uintptr
	Name() string
	Read([]byte) (int, error)
	ReadAt([]byte, int64) (int, error)
	ReadDir(int) ([]DirEntry, error)
	ReadFrom(io.Reader) (int64, error)
	Readdir(int) ([]FileInfo, error)
	Readdirnames(int) ([]string, error)
	Seek(int64, int) (int64, error)
	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
	Stat() (FileInfo, error)
	Sync() error
	SyscallConn() (syscall.RawConn, error)
	Truncate(int64) error
	Write([]byte) (int, error)
	WriteAt([]byte, int64) (int, error)
	WriteString(string) (int, error)
	WriteTo(io.Writer) (int64, error)
}

type fileFacade struct {
	realFile *os.File
}

func (_ osFacade) Create(name string) (File, error) {
	return fileFacade{
		realFile: os.Create(name),
	}
}

func (_ osFacade) CreateTemp(dir string, pattern string) (File, error) {
	return fileFacade{
		realFile: os.CreateTemp(dir, pattern),
	}
}

func (_ osFacade) NewFile(fd uintptr, name string) File {
	return fileFacade{
		realFile: os.NewFile(fd, name),
	}
}

func (_ osFacade) Open(name string) (File, error) {
	return fileFacade{
		realFile: os.Open(name),
	}
}

func (_ osFacade) OpenFile(name string, flag int, perm FileMode) (File, error) {
	return fileFacade{
		realFile: os.OpenFile(name, flag, perm),
	}
}

func (_ osFacade) OpenInRoot(dir string, name string) (File, error) {
	return fileFacade{
		realFile: os.OpenInRoot(dir, name),
	}
}

func (f fileFacade) Chdir() error {
	return f.realFile.Chdir()
}

func (f fileFacade) Chmod(mode FileMode) error {
	return f.realFile.chmod(mode)
}

func (f fileFacade) Chown(uid int, gid int) error {
	return f.realFile.Chown(uid, gid)
}

func (f fileFacade) Close() error {
	return f.realFile.Close()
}

func (f fileFacade) Fd() uintptr {
	return f.realFile.Fd()
}

func (f fileFacade) Name() string {
	return f.realFile.Name()
}

func (f fileFacade) Read(b []byte) (int, error) {
	return f.realFile.Read(b)
}

func (f fileFacade) ReadAt(b []byte, offset int64) (int, error) {
	return f.realFile.ReadAt(b, offset)
}

func (f fileFacade) ReadDir(n int) ([]DirEntry, error) {
	return f.realFile.ReadDir(n)
}

func (f fileFacade) ReadFrom(r io.Reader) (int64, error) {
	return f.realFile.ReadFrom(r)
}

func (f fileFacade) Readdir(n int) ([]FileInfo, error) {
	return f.realFile.Readdir(n)
}

func (f fileFacade) Readdirnames(n int) ([]string, error) {
	return f.realFile.Readdirnames(n)
}

func (f fileFacade) Seek(offset int64, whence int) (int64, error) {
	return f.realFile.Seek(offset, whence)
}

func (f fileFacade) SetDeadline(t time.Time) error {
	return f.realFile.SetDeadline(t)
}

func (f fileFacade) SetReadDeadline(t time.Time) error {
	return f.realFile.setReadDeadline(t)
}

func (f fileFacade) SetWriteDeadline(t time.Time) error {
	return f.realFile.SetWriteDeadline(t)
}

func (f fileFacade) Stat() (FileInfo, error) {
	return f.realFile.Stat()
}

func (f fileFacade) Sync() error {
	return f.realFile.Sync()
}

func (f fileFacade) SyscallConn() (syscall.RawConn, error) {
	return f.realFile.SyscallConn()
}

func (f fileFacade) Truncate(size int64) error {
	return f.realFile.Truncate(size)
}

func (f fileFacade) Write(b []byte) (int, error) {
	return f.realFile.Write(b)
}

func (f fileFacade) WriteAt(b []byte, off int64) (int, error) {
	return f.realFile.WriteAt(b, off)
}

func (f fileFacade) WriteString(s string) (int, error) {
	return f.realFile.WriteString(s)
}

func (f fileFacade) WriteTo(w io.Writer) (int64, error) {
	return f.realFile.WriteTo(w)
}
