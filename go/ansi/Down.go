package ansi

import "fmt"

// Down - move cursor n units
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Down(n int) *Color {
	if !disabled {
		fmt.Printf(CodeMoveDown, n)
	}
	return c
}

// Down - move cursor n units and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Down(n int) *Color {
	return (&Color{}).Down(n)
}
