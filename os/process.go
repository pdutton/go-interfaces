package os

import (
	"os"
)

type Process interface {
	PID() int

	Kill() error
	Release() error
	Signal(Signal) error
	Wait() (*ProcessState, error)
}

type processFacade struct {
	realProcess *os.Process
}

func (_ osFacade) FindProcess(pid int) (Process, error) {
	p, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}

	return processFacade{
		realProcess: p,
	}, nil
}

func (_ osFacade) StartProcess(name string, argv []string, attr *ProcAttr) (Process, error) {
	p, err := os.StartProcess(name, argv, attr)
	if err != nil {
		return nil, err
	}

	return processFacade{
		realProcess: p,
	}, nil
}

func (p processFacade) PID() int {
	return p.realProcess.Pid
}

func (p processFacade) Kill() error {
	return p.realProcess.Kill()
}

func (p processFacade) Release() error {
	return p.realProcess.Release()
}

func (p processFacade) Signal(sig Signal) error {
	return p.realProcess.Signal(sig)
}

func (p processFacade) Wait() (*ProcessState, error) {
	return p.realProcess.Wait()
}
