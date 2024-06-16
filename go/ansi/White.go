package ansi

import "fmt"

// White - set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) White() *Color {
	if !disabled {
		fmt.Print(CodeFgWhite)
	}
	return c
}

// White - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func White() *Color {
    return (&Color{}).White()
}
