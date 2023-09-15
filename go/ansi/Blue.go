package ansi

import "fmt"

// Blue - Set color
func (c *Color) Blue() *Color {
	if !disabled {
		fmt.Print(fgBlue)
	}
	return c
}

// Blue - set color and return a new color object
func Blue() *Color {
	c := Color{}
	c.Blue()
	return &c
}
