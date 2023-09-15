package ansi

import "fmt"

// White - set color
func (c *Color) White() *Color {
	if !disabled {
		fmt.Print(fgWhite)
	}
	return c
}

// White - set color and return a new color object
func White() *Color {
	c := Color{}
	c.White()
	return &c
}
