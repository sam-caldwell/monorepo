package ansi

import "fmt"

// Black - Set color
func (c *Color) Black() *Color {
	if !disabled {
		fmt.Print(fgBlack)
	}
	return c
}

// Black - set color and return a new color object
func Black() *Color {
	c := Color{}
	c.Black()
	return &c
}
