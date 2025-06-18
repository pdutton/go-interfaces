package os

import (
	"os"

	"github.com/pdutton/go-interfaces/io/fs"
)

const (
	O_RDONLY = os.O_RDONLY
	O_WRONLY = os.O_WRONLY
	O_RDWR   = os.O_RDWR

	O_APPEND = os.O_APPEND
	O_CREATE = os.O_CREATE
	O_EXCL   = os.O_EXCL
	O_SYNC   = os.O_SYNC
	O_TRUNC  = os.O_TRUNC

	SEEK_SET = os.SEEK_SET
	SEEK_CUR = os.SEEK_CUR
	SEEK_END = os.SEEK_END

	PathSeparator     = os.PathSeparator
	PathListSeparator = os.PathListSeparator

	DevNull = os.DevNull
)

var (
	ModeFile       = fs.ModeFile
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

	ModePerm = os.ModePerm
)

var (
	ErrInvalid = os.ErrInvalid

	ErrPermission = os.ErrPermission
	ErrExist      = os.ErrExist
	ErrNotExist   = os.ErrNotExist
	ErrClosed     = os.ErrClosed

	ErrNoDeadline       = os.ErrNoDeadline
	ErrDeadlineExceeded = os.ErrDeadlineExceeded

	ErrProcessDone = os.ErrProcessDone
)
