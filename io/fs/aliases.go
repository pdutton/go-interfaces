package fs

import (
    "io/fs"
)

// FSFileMode is al alias for the io/fs/FileMode type, which is
// really just an int, so octal values can be used.
type FSFileMode  = fs.FileMode

type FS          = fs.FS
type File        = fs.File
type GlobFS      = fs.GlobFS
type PathError   = fs.PathError
type ReadDirFS   = fs.ReadDirFS
type ReadDirFile = fs.ReadDirFile
type ReadFileFS  = fs.ReadFileFS
type StatFS      = fs.StatFS
type SubFS       = fs.SubFS
type WalkDirFunc = fs.WalkDirFunc

