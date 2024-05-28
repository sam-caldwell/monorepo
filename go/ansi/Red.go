package ansi

import "fmt"

// Red - Set color
func (c *Color) Red() *Color {
	if !disabled {
		fmt.Print(CodeFgRed)
	}
	return c
}

// Red - set color and return a new color object
func Red() *Color {
	c := Color{}
	c.Red()
	return &c
}
