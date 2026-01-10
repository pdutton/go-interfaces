package fs

import (
	"io/fs"
)

// The FileSystem interface provides access to all the
// functions and structs in the io/fs package.  It is
// renamed to FileSystem, instead of FS, because the
// io/fs package already contains an interface named FS.
type FileSystem interface {
	FormatDirEntry(DirEntry) string
	FormatFileInfo(FileInfo) string
	Glob(FileSystem, string) ([]string, error)
	ReadFile(FileSystem, string) ([]byte, error)
	ValidPath(string) bool
	WalkDir(FileSystem, string, WalkDirFunc) error

	// Constructors
	FileInfoToDirEntry(FileInfo) DirEntry
	ReadDir(FileSystem, string) ([]DirEntry, error)

	Sub(FileSystem, string) (FileSystem, error)

	Stat(FileSystem, string) (FileInfo, error)
}

type fileSystemFacade struct {
}

func NewFileSystem() fileSystemFacade {
	return fileSystemFacade{}
}

func (_ fileSystemFacade) FormatDirEntry(dir DirEntry) string {
	return dir.Format()
}

func (_ fileSystemFacade) FormatFileInfo(info FileInfo) string {
	return fs.FormatFileInfo(info.Nub())
}

func (_ fileSystemFacade) Glob(fsys FS, pattern string) ([]string, error) {
	return fs.Glob(fsys, pattern)
}

func (_ fileSystemFacade) ReadFile(fsys FS, name string) ([]byte, error) {
	return fs.ReadFile(fsys, name)
}

func (_ fileSystemFacade) ValidPath(name string) bool {
	return fs.ValidPath(name)
}

func (_ fileSystemFacade) WalkDir(fsys FS, root string, fn WalkDirFunc) error {
	return fs.WalkDir(fsys, root, fn)
}

func (_ fileSystemFacade) Sub(fsys FS, name string) (FS, error) {
	return fs.Sub(fsys, name)
}

func (_ fileSystemFacade) Stat(fsys FS, name string) (FileInfo, error) {
	inf, err := fs.Stat(fsys, name)
	return NewFileInfo(inf), err
}
