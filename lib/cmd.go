package hermes

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

type Command struct {
	Cmd    *exec.Cmd
	Stdout *bytes.Buffer
	Stderr *bytes.Buffer
}

type Printer struct{}

func (_ Printer) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}

func NewCommand(argv []string) *Command {
	execCmd := exec.Command(argv[0], argv[1:]...)
	cmd := &Command{
		execCmd,
		bytes.NewBuffer([]byte{}),
		bytes.NewBuffer([]byte{}),
	}

	// Tees stdout and stderr to be written both to the screen and to
	// the command instance's stdout and stderr buffers
	out, _ := execCmd.StdoutPipe()
	go io.Copy(io.MultiWriter(os.Stdout, cmd.Stdout), out)
	err, _ := execCmd.StderrPipe()
	go io.Copy(io.MultiWriter(os.Stderr, cmd.Stderr), err)

	return cmd
}

func (c *Command) Run(notifier Notifier) int {
	runErr := c.Cmd.Run()
	if runErr != nil {
		if exitErr, ok := runErr.(*exec.ExitError); ok {
			status, _ := exitErr.Sys().(syscall.WaitStatus)
			notifier.Failure(c, runErr)
			return status.ExitStatus()
		} else {
			panic(runErr)
		}
	}

	notifier.Success(c)
	return 0
}
