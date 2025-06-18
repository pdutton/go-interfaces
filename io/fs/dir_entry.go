package fs

import (
    "io/fs"
)

type DirEntry interface {
    // Access to members variables:
    Name() string
    IsDir() bool
    Type() fs.FileMode
    Info() (fs.FileInfo, error)

    // Return the underlying direntry instance
    Nub() fs.DirEntry

    format() string
}

type dirEntryFacade struct {
    nub fs.DirEntry
}

func NewDirEntry(de fs.DirEntry) dirEntryFacade {
    return dirEntryFacade{
        nub: de,
    }
}

func ToDirEntryList(in []fs.DirEntry) []DirEntry {
	var dea = []DirEntry{}

	for _, de := range in {
		dea = append(dea, NewDirEntry(de))
	}

	return dea
}

func (_ fileSystemFacade) FileInfoToDirEntry(info FileInfo) DirEntry {
    return dirEntryFacade{
        nub: fs.FileInfoToDirEntry(info),
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

func (de dirEntryFacade) Type() fs.FileMode {
    return de.nub.Type()
}

func (de dirEntryFacade) Info() (fs.FileInfo, error) {
    return de.nub.Info()
}

func (de dirEntryFacade) format() string {
    return fs.FormatDirEntry(de.nub)
}


