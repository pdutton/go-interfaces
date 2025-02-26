package exec

import (
	"context"
	"io"
	"os"
	"os/exec"
)

type Cmd interface {
	// Access to member variables:
	Path() string
	Args() []string
	Env() []string
	Dir() string
	Stdin() io.Reader
	Stdout() io.Writer
	Stderr() io.Writer
	Process() *os.Process
	ProcessState() *os.ProcessState

	// Members:
	CombinedOutput() ([]byte, error)
	Environ() []string
	Output() ([]byte, error)
	Run() error
	Start() error
	StderrPipe() (io.ReadCloser, error)
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	String() string
	Wait() error
}

type cmdFacade struct {
	realCmd *exec.Cmd

	args []string
	ctxt context.Context
}

func (_ execFacade) NewCommand(name string, options ...CommandOption) Cmd {
	var cmd = cmdFacade{}

	// Only the options that run before the real command is created
	// will run here because cmd.realCmd == nil.
	for _, f := range options {
		f(&cmd)
	}

	if cmd.ctxt == nil {
		cmd.realCmd = exec.Command(name, cmd.args...)
	} else {
		cmd.realCmd = exec.CommandContext(cmd.ctxt, name, cmd.args...)
	}

	// Only the options that run after the real command is created
	// will run here because cmd.realCmd != nil.
	for _, f := range options {
		f(&cmd)
	}

	return cmd
}

func (cmd cmdFacade) Path() string {
	return cmd.realCmd.Path
}

func (cmd cmdFacade) Args() []string {
	return cmd.realCmd.Args
}

func (cmd cmdFacade) Env() []string {
	return cmd.realCmd.Env
}

func (cmd cmdFacade) Dir() string {
	return cmd.realCmd.Dir
}

func (cmd cmdFacade) Stdin() io.Reader {
	return cmd.realCmd.Stdin
}

func (cmd cmdFacade) Stdout() io.Writer {
	return cmd.realCmd.Stdout
}

func (cmd cmdFacade) Stderr() io.Writer {
	return cmd.realCmd.Stderr
}

func (cmd cmdFacade) Process() *os.Process {
	return cmd.realCmd.Process
}

func (cmd cmdFacade) ProcessState() *os.ProcessState {
	return cmd.realCmd.ProcessState
}

func (cmd cmdFacade) CombinedOutput() ([]byte, error) {
	return cmd.realCmd.CombinedOutput()
}

func (cmd cmdFacade) Environ() []string {
	return cmd.realCmd.Environ()
}

func (cmd cmdFacade) Output() ([]byte, error) {
	return cmd.realCmd.Output()
}

func (cmd cmdFacade) Run() error {
	return cmd.realCmd.Run()
}

func (cmd cmdFacade) Start() error {
	return cmd.realCmd.Start()
}

func (cmd cmdFacade) StderrPipe() (io.ReadCloser, error) {
	return cmd.realCmd.StderrPipe()
}

func (cmd cmdFacade) StdinPipe() (io.WriteCloser, error) {
	return cmd.realCmd.StdinPipe()
}

func (cmd cmdFacade) StdoutPipe() (io.ReadCloser, error) {
	return cmd.realCmd.StdoutPipe()
}

func (cmd cmdFacade) String() string {
	return cmd.realCmd.String()
}

func (cmd cmdFacade) Wait() error {
	return cmd.realCmd.Wait()
}
