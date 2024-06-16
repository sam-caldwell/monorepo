package ansi

import "fmt"

// Magenta - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Magenta() *Color {
	if !disabled {
		fmt.Print(CodeFgMagenta)
	}
	return c
}

// Magenta - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Magenta() *Color {
	return (&Color{}).Magenta()
}
