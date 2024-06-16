package ansi

import "fmt"

// TopLeft - Move cursor to top left
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) TopLeft() *Color {
	if !disabled {
		fmt.Print(CodeSetTopLeft)
	}
	return c
}

// TopLeft - Move cursor to top left and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func TopLeft() *Color {
	return (&Color{}).TopLeft()
}
