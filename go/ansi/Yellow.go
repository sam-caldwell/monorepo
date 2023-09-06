package ansi

import "fmt"

// Yellow - set color
func (c *Color) Yellow() *Color {
	fmt.Print(fgYellow)
	return c
}

// Yellow - set color and return a new color object
func Yellow() *Color {
	c := Color{}
	c.Yellow()
	return &c
}
