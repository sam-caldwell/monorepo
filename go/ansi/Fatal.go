package ansi

import (
	"fmt"
	"os"
)

// Fatal - Terminate the program and return the exit code
func (c *Color) Fatal(exitCode int) *Color {
	fmt.Print(reset)
	os.Exit(exitCode)
	return c
}

// Fatal - Terminate the program and return the exit code
func Fatal(exitCode int) *Color {
	c := Color{}
	c.Fatal(exitCode)
	return &c
}