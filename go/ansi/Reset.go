package ansi

import (
	"fmt"
	"os"
)

// Reset - Send Reset to stdout
func (c *Color) Reset() *Color {
	defer func() { _ = os.Stdout.Sync() }()
	if !disabled {
		fmt.Print(CodeReset) // Reset color
	}
	return c
}

// Reset - Send Reset to stdout and return new color object
func Reset() *Color {
	c := Color{}
	c.Reset()
	return &c
}
