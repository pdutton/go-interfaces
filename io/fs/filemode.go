package fs

import (
    "io/fs"
)

var (
	// These "constants" are not generally useful given that
	// the FileMode interface has Is* methods, but they could
	// come in handy in unit testing so I will maintain them.

	ModeFile = NewFileMode(FSFileMode(0))

	ModeDir        = NewFileMode(fs.ModeDir)
	ModeAppend     = NewFileMode(fs.ModeAppend)
	ModeExclusive  = NewFileMode(fs.ModeExclusive)
	ModeTemporary  = NewFileMode(fs.ModeTemporary)
	ModeSymlink    = NewFileMode(fs.ModeSymlink)
	ModeDevice     = NewFileMode(fs.ModeDevice)
	ModeNamedPipe  = NewFileMode(fs.ModeNamedPipe)
	ModeSocket     = NewFileMode(fs.ModeSocket)
	ModeSetuid     = NewFileMode(fs.ModeSetuid)
	ModeSetgid     = NewFileMode(fs.ModeSetgid)
	ModeCharDevice = NewFileMode(fs.ModeCharDevice)
	ModeSticky     = NewFileMode(fs.ModeSticky)
	ModeIrregular  = NewFileMode(fs.ModeIrregular)

	// ModePerm is really just a bloated int
	ModePerm = fs.ModePerm
)

type FileMode interface {
    IsRegular() bool
    IsDir() bool
    IsAppend() bool
    IsExclusive() bool
    IsTemporary() bool
    IsSymlink() bool
    IsDevice() bool
    IsNamedPipe() bool
    IsSocket() bool
    IsSetuid() bool
    IsSetgid() bool
    IsCharDevice() bool
    IsSticky() bool
    IsIrregular() bool
    
    Perm() FSFileMode
    String() string

    // Nub retrieves the underlying implementation
    Nub() FSFileMode
}

type fileModeFacade struct {
    nub FSFileMode
}

func NewFileMode(fm FSFileMode) fileModeFacade {
    return fileModeFacade{
        nub: fm,
    }
}

func (fm fileModeFacade) Nub() FSFileMode {
	return fm.nub
}

func (fm fileModeFacade) IsRegular() bool {
	return fm.nub.IsRegular()
}

func (fm fileModeFacade) IsDir() bool {
	return fm.nub.IsDir()
}

func (fm fileModeFacade) IsAppend() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeAppend) != 0
}

func (fm fileModeFacade) IsExclusive() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeExclusive) != 0
}

func (fm fileModeFacade) IsTemporary() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeTemporary) != 0
}

func (fm fileModeFacade) IsSymlink() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeSymlink) != 0
}

func (fm fileModeFacade) IsDevice() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeDevice) != 0
}

func (fm fileModeFacade) IsNamedPipe() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeNamedPipe) != 0
}

func (fm fileModeFacade) IsSocket() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeSocket) != 0
}

func (fm fileModeFacade) IsSetuid() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeSetuid) != 0
}

func (fm fileModeFacade) IsSetgid() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeSetgid) != 0
}

func (fm fileModeFacade) IsCharDevice() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeCharDevice) != 0
}

func (fm fileModeFacade) IsSticky() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeSticky) != 0
}

func (fm fileModeFacade) IsIrregular() bool {
	return uint32(fm.nub.Type()) & uint32(fs.ModeIrregular) != 0
}

func (fm fileModeFacade) Perm() FSFileMode {
    return fm.nub.Perm()
}

func (fm fileModeFacade) String() string {
    return fm.nub.String()
}

