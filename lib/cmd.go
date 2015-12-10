package hermes

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
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

	out, _ := execCmd.StdoutPipe()
	go io.Copy(io.MultiWriter(os.Stdout, cmd.Stdout), out)
	err, _ := execCmd.StderrPipe()
	go io.Copy(io.MultiWriter(os.Stderr, cmd.Stderr), err)

	return cmd
}

func (c *Command) Run() error {
	return c.Cmd.Run()
}