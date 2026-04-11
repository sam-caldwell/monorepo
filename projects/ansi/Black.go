package ansi

import "fmt"

// Black - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Black() *Color {
	if !disabled {
		fmt.Print(CodeFgBlack)
	}
	return c
}

// Black - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Black() *Color {
	return (&Color{}).Black()
}
