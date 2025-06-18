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

	// Return the underlying process object
	Nub() *os.Process
}

type processFacade struct {
	nub *os.Process
}

func (_ osFacade) FindProcess(pid int) (Process, error) {
	p, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}

	return processFacade{
		nub: p,
	}, nil
}

func (_ osFacade) StartProcess(name string, argv []string, attr *ProcAttr) (Process, error) {
	p, err := os.StartProcess(name, argv, attr)
	if err != nil {
		return nil, err
	}

	return processFacade{
		nub: p,
	}, nil
}

func WrapProcess(p *os.Process) processFacade{
	return processFacade{
		nub: p,
	}
}

func (p processFacade) Nub() *os.Process {
	return p.nub
}

func (p processFacade) PID() int {
	return p.nub.Pid
}

func (p processFacade) Kill() error {
	return p.nub.Kill()
}

func (p processFacade) Release() error {
	return p.nub.Release()
}

func (p processFacade) Signal(sig Signal) error {
	return p.nub.Signal(sig)
}

func (p processFacade) Wait() (*ProcessState, error) {
	return p.nub.Wait()
}
