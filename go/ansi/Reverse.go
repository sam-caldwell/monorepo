package ansi

import "fmt"

// Reverse - Set format attribute
func (c *Color) Reverse() *Color {
	if !disabled {
		fmt.Print(CodeReverse)
	}
	return c
}

// Reverse - set format attribute and return a new color object
func Reverse() *Color {
	c := Color{}
	c.Reverse()
	return &c
}
