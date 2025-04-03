package filepath

import (
	"path/filepath"

	"github.com/pdutton/go-interfaces/path"
)

type FilePath interface {
	path.Path

	Abs(string) (string, error)
	EvalSymlinks(string) (string, error)
	FromSlash(string) string
	Glob(string) ([]string, error)
	IsLocal(string) bool
	Localize(string) (string, error)
	Rel(string, string) (string, error)
	SplitList(string) []string
	ToSlash(string) string
	VolumeName(string) string
	Walk(string, WalkFunc) error
	WalkDir(string, WalkDirFunc) error
}

type filePathFacade struct {
}

func NewFilePath() filePathFacade {
	return filePathFacade{}
}

func (_ filePathFacade) Base(p string) string {
	return filepath.Base(p)
}

func (_ filePathFacade) Clean(p string) string {
	return filepath.Clean(p)
}

func (_ filePathFacade) Dir(p string) string {
	return filepath.Dir(p)
}

func (_ filePathFacade) Ext(p string) string {
	return filepath.Ext(p)
}

func (_ filePathFacade) IsAbs(p string) bool {
	return filepath.IsAbs(p)
}

func (_ filePathFacade) Join(elem ...string) string {
	return filepath.Join(elem...)
}

func (_ filePathFacade) Match(pattern string, name string) (bool, error) {
	return filepath.Match(pattern, name)
}

func (_ filePathFacade) Split(p string) (string, string) {
	return filepath.Split(p)
}

func (_ filePathFacade) Abs(p string) (string, error) {
	return filepath.Abs(p)
}

func (_ filePathFacade) EvalSymlinks(p string) (string, error) {
	return filepath.EvalSymlinks(p)
}

func (_ filePathFacade) FromSlash(p string) string {
	return filepath.FromSlash(p)
}

func (_ filePathFacade) Glob(p string) ([]string, error) {
	return filepath.Glob(p)
}

func (_ filePathFacade) IsLocal(p string) bool {
	return filepath.IsLocal(p)
}

func (_ filePathFacade) Localize(p string) (string, error) {
	return filepath.Localize(p)
}

func (_ filePathFacade) Rel(basepath, targetpath string) (string, error) {
	return filepath.Rel(basepath, targetpath)
}

func (_ filePathFacade) SplitList(p string) []string {
	return filepath.SplitList(p)
}

func (_ filePathFacade) ToSlash(p string) string {
	return filepath.ToSlash(p)
}

func (_ filePathFacade) VolumeName(p string) string {
	return filepath.VolumeName(p)
}

func (_ filePathFacade) Walk(p string, fn WalkFunc) error {
	return filepath.Walk(p, fn)
}

func (_ filePathFacade) WalkDir(p string, fn WalkDirFunc) error {
	return filepath.WalkDir(p, fn)
}
