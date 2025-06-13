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
    IsDir() bool
    IsRegular() bool
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

func (fm fileModeFacade) IsDir() bool {
    return fm.realFileMode.IsDir()
}

func (fm fileModeFacade) IsRegular() bool {
    return fm.realFileMode.IsRegular()
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

