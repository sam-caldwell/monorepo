package ansi

import "fmt"

// Red - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Red() *Color {
	if !disabled {
		fmt.Print(CodeFgRed)
	}
	return c
}

// Red - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Red() *Color {
	return (&Color{}).Red()
}
