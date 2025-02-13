package os

import (
	"os"
)

const (
	O_RDONLY = os.O_RDONLY
	O_WRONLY = os.OWRONLY
	O_RDWR   = os.ORDWR

	O_APPEND = os.O_APPEND
	O_CREATE = os.O_CREATE
	O_EXCL   = os.O_EXCL
	O_SYNC   = os.SYNC
	O_TRUNC  = os.TRUNC

	SEEK_SET = os.SEEK_SET
	SEEK_CUR = os.SEEK_CUR
	SEEK_END = os.SEEK_END

	PathSeparator     = os.PathSeparator
	PathListSeparator = os.PathListSeparator

	ModeFile       = FileMode(0)
	ModeDir        = os.ModeDir
	ModeAppend     = os.ModeAppend
	ModeExclusive  = os.ModeExclusive
	ModeTemporary  = os.ModeTemporary
	ModeSymlink    = os.ModeSymlink
	ModeDevice     = os.ModeDevice
	ModeNamedPipe  = os.ModeNamedPipe
	ModeSocket     = os.ModeSocket
	ModeSetuid     = os.ModeSetuid
	ModeSetgid     = os.ModeSetgid
	ModeCharDevice = os.ModeCharDevice
	ModeSticky     = os.ModeSticky
	ModeIrregular  = os.ModeIrregular

	ModeType = os.ModeType

	ModePerm = os.ModePerm

	DevNull = os.DevNull
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
