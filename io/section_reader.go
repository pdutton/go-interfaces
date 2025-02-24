package io

import (
	"io"
)

type SectionReader interface {
	Outer() (ReaderAt, int64, int64)
	Read([]byte) (int, error)
	ReadAt([]byte, int64) (int, error)
	Seek(int64, int) (int64, error)
	Size() int64
}

type sectionReaderFacade struct {
	realSectionReader *io.SectionReader
}

func (_ ioFacade) NewSectionReader(r ReaderAt, off int64, n int64) SectionReader {
	return sectionReaderFacade{
		realSectionReader: io.NewSectionReader(r, off, n),
	}
}

func (sr sectionReaderFacade) Outer() (ReaderAt, int64, int64) {
	return sr.realSectionReader.Outer()
}

func (sr sectionReaderFacade) Read(p []byte) (int, error) {
	return sr.realSectionReader.Read(p)
}

func (sr sectionReaderFacade) ReadAt(p []byte, off int64) (int, error) {
	return sr.realSectionReader.ReadAt(p, off)
}

func (sr sectionReaderFacade) Seek(offset int64, whence int) (int64, error) {
	return sr.realSectionReader.Seek(offset, whence)
}

func (sr sectionReaderFacade) Size() int64 {
	return sr.realSectionReader.Size()
}
