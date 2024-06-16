package ansi

import "fmt"

// Up - move cursor n units
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Up(n int) *Color {
	if !disabled {
		fmt.Printf(CodeMoveUp, n)
	}
	return c
}

// Up - move cursor n units and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Up(n int) *Color {
	return (&Color{}).Up(n)
}
