package tmux

import (
	"bytes"
	"os"
	"os/exec"
	"syscall"
)

// Command abstracts tmux command
type Command struct {
	exe  string
	args []string
}

// NewCommand create instance of Command
func NewCommand(args ...string) Command {
	return Command{
		exe:  "tmux",
		args: append([]string{"-2"}, args...),
	}
}

// Run executes tmux command
func (c Command) Run() ([]byte, error) {
	cmd := exec.Command(c.exe, c.args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return bytes.TrimSpace(stdout.Bytes()), nil
}

// Exec command replacing current process
func (c Command) Exec() error {
	exePath, err := exec.LookPath(c.exe)
	if err != nil {
		return err
	}
	return syscall.Exec(exePath, append([]string{c.exe}, c.args...), os.Environ())
}
