package ansi

import "fmt"

// Underline - Set format attribute
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Underline() *Color {
	if !disabled {
		fmt.Print(CodeUnderline)
	}
	return c
}

// Underline - set format attribute and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Underline() *Color {
	return (&Color{}).Underline()
}
