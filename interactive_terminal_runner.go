package interactive_terminal_runner

import (
	"fmt"
	"os"
	"os/exec"
)

type TerminalRunnerOptsFunc func(runner *terminalRunner) error

type terminalRunner struct {
	execPath string
}

// NewTerminalRunner takes customer runner config and returns a terminal runner or an error
func NewTerminalRunner(cfgs ...TerminalRunnerOptsFunc) (*terminalRunner, error) {
	path, err := exec.LookPath("bash")

	if err != nil {
		return nil, err
	}

	tr := &terminalRunner{
		execPath: path,
	}

	for _, fn := range cfgs {
		if err = fn(tr); err != nil {
			return nil, err
		}
	}

	return tr, nil
}

// Exec takes the programme you wish to run and its args and executes the programme
// an error is returned if an error occurs
// example programme = docker
// args are what is passed to the programme
func (r *terminalRunner) Exec(programme string, args string) error {
	programmeExec, err := r.getExecPath(programme)

	if err != nil {
		return err
	}

	c := exec.Command(r.execPath, "-c", fmt.Sprintf("%s %s", programmeExec, args))

	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err = c.Run(); err != nil {
		return err
	}

	sttyExec, err := r.getExecPath("/bin/stty")

	if err != nil {
		return err
	}

	c = exec.Command(r.execPath, "-c", fmt.Sprintf("%s %s", sttyExec, "sane"))
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err = c.Run(); err != nil {
		return err
	}

	return nil
}

// WithCustomExecPath takes a programme such as sh and will run the entire process via that instead
// Defaults to bash
func WithCustomExecPath(programme string) TerminalRunnerOptsFunc {
	return func(r *terminalRunner) error {
		execPath, err := r.getExecPath(programme)

		if err != nil {
			return err
		}

		r.execPath = execPath

		return nil
	}
}

func (r *terminalRunner) getExecPath(programme string) (string, error) {
	path, err := exec.LookPath(programme)

	if err != nil {
		return "", err
	}

	return path, nil
}
