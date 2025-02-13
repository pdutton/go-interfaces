package os

import (
	"io/fs"
	"os"
)

type Root interface {
	Close() error
	Create(string) (File, error)
	FS() fs.FS
	Lstat(string) (FileInfo, error)
	Mkdir(string, FileMode) error
	Name() string
	Open(string) (File, error)
	OpenFile(string, int, FileMode) (File, error)
	OpenRoot(string) (Root, error)
	Remove(string) error
	Stat(string) (FileInfo, error)
}

type rootFacade struct {
	realRoot *os.Root
}

func (_ osFacade) OpenRoot(name string) (Root, error) {
	r, err := os.OpenRoot(name)
	if err != nil {
		return nil, err
	}

	return rootFacade{
		realRoot: r,
	}, nil
}

func (r rootFacade) Close() error {
	return r.realRoot.Close()
}

func (r rootFacade) Create(name string) (File, error) {
	f, err := r.realRoot.Create(name)
	if err != nil {
		return nil, err
	}

	return fileFacade{
		realFile: f,
	}, nil
}

func (r rootFacade) FS() fs.FS {
	return r.realRoot.FS()
}

func (r rootFacade) Lstat(name string) (FileInfo, error) {
	return r.realRoot.Lstat(name)
}

func (r rootFacade) Mkdir(name string, perm FileMode) error {
	return r.realRoot.Mkdir(name, perm)
}

func (r rootFacade) Name() string {
	return r.realRoot.Name()
}

func (r rootFacade) Open(name string) (File, error) {
	f, err := r.realRoot.Open(name)
	if err != nil {
		return nil, err
	}

	return fileFacade{
		realFile: f,
	}, nil
}

func (r rootFacade) OpenFile(name string, flag int, perm FileMode) (File, error) {
	return r.realRoot.OpenFile(name, flag, perm)
}

func (r rootFacade) OpenRoot(name string) (Root, error) {
	r2, err := r.realRoot.OpenRoot(name)
	if err != nil {
		return nil, err
	}

	return rootFacade{
		realRoot: r2,
	}, nil
}

func (r rootFacade) Remove(name string) error {
	return r.realRoot.Remove(name)
}

func (r rootFacade) Stat(name string) (FileInfo, error) {
	return r.realRoot.Stat(name)
}
