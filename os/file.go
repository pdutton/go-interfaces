package os

import (
	"io"
	rfs "io/fs"
	"os"
	"syscall"
	"time"

	"github.com/pdutton/go-interfaces/io/fs"
)

type File interface {
	Chdir() error
	Chmod(rfs.FileMode) error
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
	nub *os.File
}

// This is a basic constructor, but I didn't want to name it New*
// because that might imply that it actually creates a file.
func WrapFile(f *os.File) fileFacade {
	return fileFacade{
		nub: f,
	}
}

func (_ osFacade) Create(name string) (File, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return fileFacade{
		nub: f,
	}, nil
}

func (_ osFacade) CreateTemp(dir string, pattern string) (File, error) {
	f, err := os.CreateTemp(dir, pattern)
	if err != nil {
		return nil, err
	}

	return fileFacade{
		nub: f,
	}, nil
}

func (_ osFacade) NewFile(fd uintptr, name string) File {
	return fileFacade{
		nub: os.NewFile(fd, name),
	}
}

func (_ osFacade) Open(name string) (File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return fileFacade{
		nub: f,
	}, nil
}

func (_ osFacade) OpenFile(name string, flag int, perm rfs.FileMode) (File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return fileFacade{
		nub: f,
	}, nil
}

func (_ osFacade) OpenInRoot(dir string, name string) (File, error) {
	f, err := os.OpenInRoot(dir, name)
	if err != nil {
		return nil, err
	}

	return fileFacade{
		nub: f,
	}, nil
}

func (f fileFacade) Nub() *os.File {
	return f.nub
}

func (f fileFacade) Chdir() error {
	return f.nub.Chdir()
}

func (f fileFacade) Chmod(mode rfs.FileMode) error {
	return f.nub.Chmod(mode)
}

func (f fileFacade) Chown(uid int, gid int) error {
	return f.nub.Chown(uid, gid)
}

func (f fileFacade) Close() error {
	return f.nub.Close()
}

func (f fileFacade) Fd() uintptr {
	return f.nub.Fd()
}

func (f fileFacade) Name() string {
	return f.nub.Name()
}

func (f fileFacade) Read(b []byte) (int, error) {
	return f.nub.Read(b)
}

func (f fileFacade) ReadAt(b []byte, offset int64) (int, error) {
	return f.nub.ReadAt(b, offset)
}

func (f fileFacade) ReadDir(n int) ([]DirEntry, error) {
	dea, err := f.nub.ReadDir(n)
	return fs.NewDirEntryList(dea), err
}

func (f fileFacade) ReadFrom(r io.Reader) (int64, error) {
	return f.nub.ReadFrom(r)
}

func (f fileFacade) Readdir(n int) ([]FileInfo, error) {
	fil, err := f.nub.Readdir(n)
	return fs.NewFileInfoList(fil), err
}

func (f fileFacade) Readdirnames(n int) ([]string, error) {
	return f.nub.Readdirnames(n)
}

func (f fileFacade) Seek(offset int64, whence int) (int64, error) {
	return f.nub.Seek(offset, whence)
}

func (f fileFacade) SetDeadline(t time.Time) error {
	return f.nub.SetDeadline(t)
}

func (f fileFacade) SetReadDeadline(t time.Time) error {
	return f.nub.SetReadDeadline(t)
}

func (f fileFacade) SetWriteDeadline(t time.Time) error {
	return f.nub.SetWriteDeadline(t)
}

func (f fileFacade) Stat() (FileInfo, error) {
	fi, err := f.nub.Stat()
	return fs.NewFileInfo(fi), err
}

func (f fileFacade) Sync() error {
	return f.nub.Sync()
}

func (f fileFacade) SyscallConn() (syscall.RawConn, error) {
	return f.nub.SyscallConn()
}

func (f fileFacade) Truncate(size int64) error {
	return f.nub.Truncate(size)
}

func (f fileFacade) Write(b []byte) (int, error) {
	return f.nub.Write(b)
}

func (f fileFacade) WriteAt(b []byte, off int64) (int, error) {
	return f.nub.WriteAt(b, off)
}

func (f fileFacade) WriteString(s string) (int, error) {
	return f.nub.WriteString(s)
}

func (f fileFacade) WriteTo(w io.Writer) (int64, error) {
	return f.nub.WriteTo(w)
}
