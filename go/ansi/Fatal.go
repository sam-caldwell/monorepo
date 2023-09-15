package ansi

import (
	"os"
)

// Fatal - Terminate the program and return the exit code
func (c *Color) Fatal(exitCode int) *Color {
	Reset()
	os.Exit(exitCode)
	return c
}

// Fatal - Terminate the program and return the exit code
func Fatal(exitCode int) *Color {
	c := Color{}
	c.Fatal(exitCode)
	return &c
}
