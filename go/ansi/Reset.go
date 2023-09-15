package ansi

import "fmt"

// Reset - Send Reset to stdout
func (c *Color) Reset() *Color {
	if !disabled {
		fmt.Print(reset) // Reset color
	}
	return c
}

// Reset - Send Reset to stdout and return new color object
func Reset() *Color {
	c := Color{}
	c.Reset()
	return &c
}
