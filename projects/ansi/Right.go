package ansi

import "fmt"

// Right - move cursor n units
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Right(n int) *Color {
	if !disabled {
		fmt.Printf(CodeMoveRight, n)
	}
	return c
}

// Right - move cursor n units and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Right(n int) *Color {
	return (&Color{}).Right(n)
}
