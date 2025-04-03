package fs

import (
    "io/fs"
)

var (
    ErrInvalid    = fs.ErrInvalid
    ErrPermission = fs.ErrPermission
    ErrExist      = fs.ErrExist
    ErrNotExist   = fs.ErrNotExist
    ErrClosed     = fs.ErrClosed

    SkipAll = fs.SkipAll
    SkipDir = fs.SkipDir
)
