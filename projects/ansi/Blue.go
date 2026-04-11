package ansi

import "fmt"

// Blue - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Blue() *Color {
	if !disabled {
		fmt.Print(CodeFgBlue)
	}
	return c
}

// Blue - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Blue() *Color {
	return (&Color{}).Blue()
}
