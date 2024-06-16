package ansi

import "fmt"

// Strikethrough - Set format attribute
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Strikethrough() *Color {
	if !disabled {
		fmt.Print(CodeStrikeThrough)
	}
	return c
}

// Strikethrough - set format attribute and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Strikethrough() *Color {
	return (&Color{}).Strikethrough()
}
