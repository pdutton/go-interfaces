package os

import (
	"os"
	"time"
)

type ProcessState interface {
	ExitCode() int
	Exited() bool
	Pid() int
	String() string
	Success() bool
	Sys() any
	SysUsage() any
	SystemTime() time.Duration
	UserTime() time.Duration
	Nub() *os.ProcessState
}

type processStateFacade struct {
	nub *os.ProcessState
}

func WrapProcessState(ps *os.ProcessState) processStateFacade {
	return processStateFacade{nub: ps}
}

func (ps processStateFacade) Nub() *os.ProcessState {
	return ps.nub
}

func (ps processStateFacade) ExitCode() int {
	return ps.nub.ExitCode()
}

func (ps processStateFacade) Exited() bool {
	return ps.nub.Exited()
}

func (ps processStateFacade) Pid() int {
	return ps.nub.Pid()
}

func (ps processStateFacade) String() string {
	return ps.nub.String()
}

func (ps processStateFacade) Success() bool {
	return ps.nub.Success()
}

func (ps processStateFacade) Sys() any {
	return ps.nub.Sys()
}

func (ps processStateFacade) SysUsage() any {
	return ps.nub.SysUsage()
}

func (ps processStateFacade) SystemTime() time.Duration {
	return ps.nub.SystemTime()
}

func (ps processStateFacade) UserTime() time.Duration {
	return ps.nub.UserTime()
}
