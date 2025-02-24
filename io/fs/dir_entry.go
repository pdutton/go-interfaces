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

    format() string
}

type dirEntryFacade struct {
    realDirEntry fs.DirEntry
}

func newDirEntry(de fs.DirEntry) dirEntryFacade {
    return dirEntryFacade{
        realDirEntry: de,
    }
}

func (_ fileSystemFacade) FileInfoToDirEntry(info FileInfo) DirEntry {
    return dirEntryFacade{
        realDirEntry: fs.FileInfoToDirEntry(info),
    }
}

func (_ fileSystemFacade) ReadDir(fsys FS, name string) ([]DirEntry, error) {
    entries, err := fs.ReadDir(fsys, name)
    if err != nil {
        return nil, err
    }

    var results = []DirEntry{}
    for _, entry := range entries {
        results = append(results, dirEntryFacade{ realDirEntry: entry})
    }

    return results, nil
}

func (de dirEntryFacade) Name() string {
    return de.realDirEntry.Name()
}

func (de dirEntryFacade) IsDir() bool {
    return de.realDirEntry.IsDir()
}

func (de dirEntryFacade) Type() fs.FileMode {
    return de.realDirEntry.Type()
}

func (de dirEntryFacade) Info() (fs.FileInfo, error) {
    return de.realDirEntry.Info()
}

func (de dirEntryFacade) format() string {
    return fs.FormatDirEntry(de.realDirEntry)
}


