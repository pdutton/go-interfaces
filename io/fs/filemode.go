package fs

import (
    "io/fs"
)

const (
    ModeRegular fs.FileMode = 0

    ModeDir        = fs.ModeDir
    ModeAppend     = fs.ModeAppend
    ModeExclusive  = fs.ModeExclusive
    ModeTemporary  = fs.ModeTemporary
    ModeSymlink    = fs.ModeSymlink
    ModeDevice     = fs.ModeDevice
    ModeNamedPipe  = fs.ModeNamedPipe
    ModeSocket     = fs.ModeSocket
    ModeSetuid     = fs.ModeSetuid
    ModeSetgid     = fs.ModeSetgid
    ModeCharDevice = fs.ModeCharDevice
    ModeSticky     = fs.ModeSticky
    ModeIrregular  = fs.ModeIrregular

    ModeType = fs.ModeType
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
    
    Perm() FileMode
    String() string
    Type() FileMode
}

type fileModeFacade struct {
    realFileMode fs.FileMode
}

func newFileMode(fm fs.FileMode) fileModeFacade {
    return fileModeFacade{
        realFileMode: fm,
    }
}

func (fm fileModeFacade) IsRegular() bool {
	return fm.realFileMode.IsRegular()
}

func (fm fileModeFacade) IsDir() bool {
	return fm.realFileMode.IsDir()
}

func (fm fileModeFacade) IsAppend() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeAppend) != 0
}

func (fm fileModeFacade) IsExclusive() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeExclusive) != 0
}

func (fm fileModeFacade) IsTemporary() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeTemporary) != 0
}

func (fm fileModeFacade) IsSymlink() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeSymlink) != 0
}

func (fm fileModeFacade) IsDevice() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeDevice) != 0
}

func (fm fileModeFacade) IsNamedPipe() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeNamedPipe) != 0
}

func (fm fileModeFacade) IsSocket() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeSocket) != 0
}

func (fm fileModeFacade) IsSetuid() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeSetuid) != 0
}

func (fm fileModeFacade) IsSetgid() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeSetgid) != 0
}

func (fm fileModeFacade) IsCharDevice() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeCharDevice) != 0
}

func (fm fileModeFacade) IsSticky() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeSticky) != 0
}

func (fm fileModeFacade) IsIrregular() bool {
	return uint32(fm.realFileMode.Type()) & uint32(ModeIrregular) != 0
}

func (fm fileModeFacade) Perm() FileMode {
    return newFileMode(fm.realFileMode.Perm())
}

func (fm fileModeFacade) String() string {
    return fm.realFileMode.String()
}

func (fm fileModeFacade) Type() FileMode {
    return newFileMode(fm.realFileMode.Type())
}

