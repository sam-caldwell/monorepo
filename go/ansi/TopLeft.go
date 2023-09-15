package ansi

import "fmt"

// TopLeft - Move cursor to top left
func (c *Color) TopLeft() *Color {
	if !disabled {
		fmt.Print(setTopLeft)
	}
	return c
}

// TopLeft - Move cursor to top left and return a new color object
func TopLeft() *Color {
	c := Color{}
	c.TopLeft()
	return &c
}
