package util

import (
	"bytes"
	"context"
	"os/exec"
	"time"
)

// Command ...
type Command struct {
	Cmd     []string
	stdout  []byte
	stderr  []byte
	Timeout time.Duration
	failed  bool
}

// Run ...
func (c *Command) Run() {
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	ctx, cancel := context.WithTimeout(context.Background(), c.Timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, c.Cmd[0], c.Cmd[1:]...)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	startErr := cmd.Start()
	startFailed := (startErr != nil)
	waitErr := cmd.Wait()
	waitFailed := (waitErr != nil)
	c.failed = startFailed || waitFailed || false
	c.stdout = stdoutBuf.Bytes()
	c.stderr = stderrBuf.Bytes()
}

// Stdout ...
func (c *Command) Stdout() []byte {
	return c.stdout
}

// Stderr ...
func (c *Command) Stderr() []byte {
	return c.stderr
}

// Failed ...
func (c *Command) Failed() bool {
	return c.failed
}
