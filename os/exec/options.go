package exec

import (
	"context"
	"io"
	"os"
	"syscall"
	"time"
)

type CommandOption func(*cmdFacade)

func WithArgs(args ...string) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd == nil {
			cmd.args = args
		}
	}
}

func WithContext(ctxt context.Context) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd == nil {
			cmd.ctxt = ctxt
		}
	}
}

func WithEnv(name, value string) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd != nil {
			cmd.realCmd.Env = append(cmd.realCmd.Env, name+`=`+value)
		}
	}
}

func WithDir(dir string) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd != nil {
			cmd.realCmd.Dir = dir
		}
	}
}

func WithStdin(r io.Reader) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd != nil {
			cmd.realCmd.Stdin = r
		}
	}
}

func WithStdout(w io.Writer) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd != nil {
			cmd.realCmd.Stdout = w
		}
	}
}

func WithStderr(w io.Writer) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd != nil {
			cmd.realCmd.Stderr = w
		}
	}
}

func WithExtraFiles(files []*os.File) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd != nil {
			cmd.realCmd.ExtraFiles = files
		}
	}
}

func WithSysProcAttr(spa *syscall.SysProcAttr) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd != nil {
			cmd.realCmd.SysProcAttr = spa
		}
	}
}

func WithCancel(f func() error) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd != nil {
			cmd.realCmd.Cancel = f
		}
	}
}

func WithWaitDelay(d time.Duration) CommandOption {
	return func(cmd *cmdFacade) {
		if cmd.realCmd != nil {
			cmd.realCmd.WaitDelay = d
		}
	}
}
