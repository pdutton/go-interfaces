package path

import (
	"path"
)

// These functions are also aliased so you can use them directly
// if you want to mock some but not others.

var (
	Base  = path.Base
	Clean = path.Clean
	Dir   = path.Dir
	Ext   = path.Ext
	IsAbs = path.IsAbs
	Join  = path.Join
	Match = path.Match
	Split = path.Split
)
