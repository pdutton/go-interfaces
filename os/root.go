package os

import (
	"os"

	"github.com/pdutton/go-interfaces/io/fs"
)

type Root interface {
	Close() error
	Create(string) (File, error)
	FS() fs.FS
	Lstat(string) (FileInfo, error)
	Mkdir(string, FSFileMode) error
	Name() string
	Open(string) (File, error)
	OpenFile(string, int, FSFileMode) (File, error)
	OpenRoot(string) (Root, error)
	Remove(string) error
	Stat(string) (FileInfo, error)

	Nub() *os.Root
}

type rootFacade struct {
	nub *os.Root
}

func (_ osFacade) OpenRoot(name string) (Root, error) {
	r, err := os.OpenRoot(name)
	if err != nil {
		return nil, err
	}

	return rootFacade{
		nub: r,
	}, nil
}

func (r rootFacade) Nub() *os.Root {
	return r.nub
}

func (r rootFacade) Close() error {
	return r.nub.Close()
}

func (r rootFacade) Create(name string) (File, error) {
	f, err := r.nub.Create(name)
	if err != nil {
		return nil, err
	}

	return WrapFile(f), nil
}

func (r rootFacade) FS() fs.FS {
	return r.nub.FS()
}

func (r rootFacade) Lstat(name string) (FileInfo, error) {
	fi, err := r.nub.Lstat(name)
	return fs.NewFileInfo(fi), err
}

func (r rootFacade) Mkdir(name string, perm FSFileMode) error {
	return r.nub.Mkdir(name, perm)
}

func (r rootFacade) Name() string {
	return r.nub.Name()
}

func (r rootFacade) Open(name string) (File, error) {
	f, err := r.nub.Open(name)
	if err != nil {
		return nil, err
	}

	return WrapFile(f), nil
}

func (r rootFacade) OpenFile(name string, flag int, perm FSFileMode) (File, error) {
	f, err := r.nub.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return WrapFile(f), nil
}

func (r rootFacade) OpenRoot(name string) (Root, error) {
	r2, err := r.nub.OpenRoot(name)
	if err != nil {
		return nil, err
	}

	return rootFacade{
		nub: r2,
	}, nil
}

func (r rootFacade) Remove(name string) error {
	return r.nub.Remove(name)
}

func (r rootFacade) Stat(name string) (FileInfo, error) {
	fi, err := r.nub.Stat(name)
	return fs.NewFileInfo(fi), err
}
