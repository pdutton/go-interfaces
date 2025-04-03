package path

import (
	"path"
)

type Path interface {
	Base(string) string
	Clean(string) string
	Dir(string) string
	Ext(string) string
	IsAbs(string) bool
	Join(...string) string
	Match(string, string) (bool, error)
	Split(string) (string, string)
}

type pathFacade struct {
}

func NewPath() pathFacade {
	return pathFacade{}
}

func (_ pathFacade) Base(p string) string {
	return path.Base(p)
}

func (_ pathFacade) Clean(p string) string {
	return path.Clean(p)
}

func (_ pathFacade) Dir(p string) string {
	return path.Dir(p)
}

func (_ pathFacade) Ext(p string) string {
	return path.Ext(p)
}

func (_ pathFacade) IsAbs(p string) bool {
	return path.IsAbs(p)
}

func (_ pathFacade) Join(elem ...string) string {
	return path.Join(elem...)
}

func (_ pathFacade) Match(pattern string, name string) (bool, error) {
	return path.Match(pattern, name)
}

func (_ pathFacade) Split(p string) (string, string) {
	return path.Split(p)
}
