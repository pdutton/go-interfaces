package filepath

import (
	"path/filepath"
)

// These functions are also aliased so you can use them directly
// if you want to mock some but not others.

var (
	Abs          = filepath.Abs
	Base         = filepath.Base
	Clean        = filepath.Clean
	Dir          = filepath.Dir
	EvalSymlinks = filepath.EvalSymlinks
	Ext          = filepath.Ext
	FromSlash    = filepath.FromSlash
	Glob         = filepath.Glob
	IsAbs        = filepath.IsAbs
	IsLocal      = filepath.IsLocal
	Join         = filepath.Join
	Localize     = filepath.Localize
	Match        = filepath.Match
	Rel          = filepath.Rel
	Split        = filepath.Split
	SplitList    = filepath.SplitList
	ToSlash      = filepath.ToSlash
	VolumeName   = filepath.VolumeName
	Walk         = filepath.Walk
	WalkDir      = filepath.WalkDir
)
