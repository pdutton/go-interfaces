package os

import (
	"os"

	"github.com/pdutton/go-interfaces/io/fs"
)

type DirEntry = fs.DirEntry
type FileInfo = fs.FileInfo
type FileMode = fs.FileMode

type LinkError = os.LinkError
type PathError = os.PathError
type ProcAttr = os.ProcAttr
type ProcessState = os.ProcessState
type Signal = os.Signal
type SyscallError = os.SyscallError
