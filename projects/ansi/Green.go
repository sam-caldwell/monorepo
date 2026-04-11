package ansi

import "fmt"

// Green - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Green() *Color {
	if !disabled {
		fmt.Print(CodeFgGreen)
	}
	return c
}

// Green - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Green() *Color {
	return (&Color{}).Green()
}
