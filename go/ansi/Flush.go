package ansi

import (
	"os"
)

// Flush - Flush StdOut buffer
func (c *Color) Flush() *Color {
	defer func() { _ = os.Stdout.Sync() }()
	return c
}

// Flush - Flush StdOut buffer
func Flush() *Color {
	c := Color{}
	c.Flush()
	return &c
}
