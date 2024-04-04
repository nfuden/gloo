package cmdutils

import (
	"context"
	"io"
)

// Cmd abstracts over running a command somewhere
type Cmd interface {
	// Run executes the command (like os/exec.Cmd.Run)
	// It returns a *RunError if there is any error, nil otherwise
	Run() *RunError

	// SetEnv sets the Env variables for the Cmd
	// Each entry should be of the form "key=value"
	SetEnv(...string) Cmd

	// SetStdin sets the io.Reader used for stdin
	SetStdin(reader io.Reader) Cmd

	// SetStdout sets the io.Writer used for stdout
	SetStdout(io.Writer) Cmd

	// SetStderr sets the io.Reader used for stderr
	SetStderr(io.Writer) Cmd
}

// Cmder abstracts over creating commands
type Cmder interface {
	Command(ctx context.Context, name string, args ...string) Cmd
}