package ansi

import "fmt"

// Magenta - Set color
func (c *Color) Magenta() *Color {
	if !disabled {
		fmt.Print(fgMagenta)
	}
	return c
}

// Magenta - set color and return a new color object
func Magenta() *Color {
	c := Color{}
	c.Magenta()
	return &c
}
