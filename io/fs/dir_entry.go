package fs

import (
    "io/fs"
)

type DirEntry interface {
    // Access to members variables:
    Name() string
    IsDir() bool
    Type() FileMode
    Info() (FileInfo, error)

    // Return the underlying direntry instance
    Nub() fs.DirEntry

    Format() string
}

type dirEntryFacade struct {
    nub fs.DirEntry
}

func NewDirEntry(de fs.DirEntry) dirEntryFacade {
    return dirEntryFacade{
        nub: de,
    }
}

func NewDirEntryList(in []fs.DirEntry) []DirEntry {
	var dea = []DirEntry{}

	for _, de := range in {
		dea = append(dea, NewDirEntry(de))
	}

	return dea
}

func (_ fileSystemFacade) FileInfoToDirEntry(info FileInfo) DirEntry {
    return dirEntryFacade{
        nub: fs.FileInfoToDirEntry(info.Nub()),
    }
}

func (_ fileSystemFacade) ReadDir(fsys FS, name string) ([]DirEntry, error) {
    entries, err := fs.ReadDir(fsys, name)
    if err != nil {
        return nil, err
    }

    var results = []DirEntry{}
    for _, entry := range entries {
        results = append(results, dirEntryFacade{ nub: entry})
    }

    return results, nil
}

func (de dirEntryFacade) Nub() fs.DirEntry {
	return de.nub
}

func (de dirEntryFacade) Name() string {
    return de.nub.Name()
}

func (de dirEntryFacade) IsDir() bool {
    return de.nub.IsDir()
}

func (de dirEntryFacade) Type() FileMode {
    return NewFileMode(de.nub.Type())
}

func (de dirEntryFacade) Info() (FileInfo, error) {
	inf, err := de.nub.Info()
	return NewFileInfo(inf), err
}

func (de dirEntryFacade) Format() string {
    return fs.FormatDirEntry(de.nub)
}


