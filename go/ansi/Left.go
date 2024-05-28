package ansi

import "fmt"

// Left - move cursor n units
func (c *Color) Left(n int) *Color {
	if !disabled {
		fmt.Printf(CodeMoveLeft, n)
	}
	return c
}

// Left - move cursor n units and return a new color object
func Left(n int) *Color {
	c := Color{}
	c.Left(n)
	return &c
}
