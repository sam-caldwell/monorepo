package ansi

import "fmt"

// Reverse - Set format attribute
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Reverse() *Color {
	if !disabled {
		fmt.Print(CodeReverse)
	}
	return c
}

// Reverse - set format attribute and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Reverse() *Color {
	return (&Color{}).Reverse()
}
