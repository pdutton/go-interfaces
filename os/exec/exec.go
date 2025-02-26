package exec

import (
	"os/exec"
)

type Exec interface {
	// Functions:
	LookPath(string) (string, error)

	// Constructors:
	NewCommand(string, ...CommandOption) Cmd
}

type execFacade struct {
}

func NewExec() execFacade {
	return execFacade{}
}

func (_ execFacade) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}
