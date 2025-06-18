package fs

import (
	"io/fs"
	"time"
)

type FileInfo interface {
	Name() string
	Size() int64
	Mode() FileMode
	ModTime() time.Time
	IsDir() bool
	Sys() any

	Nub() fs.FileInfo
}

type fileInfoFacade struct {
	nub fs.FileInfo
}

func NewFileInfo(fi fs.FileInfo) fileInfoFacade {
	return fileInfoFacade{
		nub: fi,
	}
}

func NewFileInfoList(fil []fs.FileInfo) []FileInfo {
	var fia = []FileInfo{}

	for _, fi := range fil {
		fia = append(fia, NewFileInfo(fi))
	}

	return fia 
}

func (fi fileInfoFacade) Nub() fs.FileInfo {
	return fi.nub
}

func (fi fileInfoFacade) Name() string {
	return fi.nub.Name()
}

func (fi fileInfoFacade) Size() int64 {
	return fi.nub.Size()
}

func (fi fileInfoFacade) Mode() FileMode {
	return NewFileMode(fi.nub.Mode())
}

func (fi fileInfoFacade) ModTime() time.Time {
	return fi.nub.ModTime()
}

func (fi fileInfoFacade) IsDir() bool {
	return fi.nub.IsDir()
}

func (fi fileInfoFacade) Sys() any {
	return fi.nub.Sys()
}



