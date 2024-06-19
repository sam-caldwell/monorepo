package ansi

import (
	"github.com/sam-caldwell/monorepo/go/exit/safety"
	"os"
)

// Fatal - Terminate the program and return the exit code
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Fatal(exitCode int) *Color {
	if !disabled {
		Reset()
	}
	safety.Defer(func() { os.Exit(exitCode) })
	return c
}

// Fatal - Terminate the program and return the exit code
//
//	(c) 2023 Sam Caldwell.  MIT License
func Fatal(exitCode int) *Color {
	return (&Color{}).Fatal(exitCode)
}
