package ansi

import (
	"fmt"
	"os"
)

// Reset - Send Reset to stdout
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Reset() *Color {
	defer func() { _ = os.Stdout.Sync() }()
	if !disabled {
		fmt.Print(CodeReset) // Reset color
	}
	return c
}

// Reset - Send Reset to stdout and return new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Reset() *Color {
	return (&Color{}).Red()
}
