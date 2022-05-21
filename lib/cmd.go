package hermes

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	Cmd    *exec.Cmd
	Stdout *bytes.Buffer
	Stderr *bytes.Buffer
}

func peek(reader io.Reader, numlines int) string {
	var lines = []string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) == numlines {
			break
		}
	}
	return strings.Join(lines, "\n")
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

func (c *Command) CmdLine() string {
	return strings.Join(c.Cmd.Args, " ")
}

func (c *Command) PeekStdout(numlines int) string {
	return peek(c.Stdout, numlines)
}

func (c *Command) PeekStderr(numlines int) string {
	return peek(c.Stderr, numlines)
}

func (c *Command) ExitCode() int {
	return c.Cmd.ProcessState.ExitCode()
}

func (c *Command) Successful() bool {
	return c.ExitCode() == 0
}

func (c *Command) Run(notifier Notifier) int {
	runErr := c.Cmd.Run()

	if runErr != nil {
		if _, ok := runErr.(*exec.ExitError); !ok {
			panic(runErr)
		}
	}

	if notifyErr := notifier.Notify(c); notifyErr != nil {
		panic(notifyErr)
	}
	return c.ExitCode()
}
