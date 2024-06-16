package ansi

import "fmt"

// Left - move cursor n units
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Left(n int) *Color {
	if !disabled {
		fmt.Printf(CodeMoveLeft, n)
	}
	return c
}

// Left - move cursor n units and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Left(n int) *Color {
	return (&Color{}).Left(n)
}
