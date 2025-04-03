package filepath

import (
	"io/fs"
	"path/filepath"
)

type DirEntry = fs.DirEntry
type FileInfo = fs.FileInfo
type WalkDirFunc = fs.WalkDirFunc
type WalkFunc = filepath.WalkFunc
