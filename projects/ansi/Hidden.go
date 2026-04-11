package ansi

import "fmt"

// Hidden - Set format attribute
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Hidden() *Color {
	if !disabled {
		fmt.Print(CodeHiddenText)
	}
	return c
}

// Hidden - set format attribute and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Hidden() *Color {
	return (&Color{}).Hidden()
}
