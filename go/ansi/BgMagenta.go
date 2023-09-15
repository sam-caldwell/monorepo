package ansi

import "fmt"

// BgMagenta - Set color
func (c *Color) BgMagenta() *Color {
	if !disabled {
		fmt.Print(bgMagenta)
	}
	return c
}

// BgMagenta - set color and return a new color object
func BgMagenta() *Color {
	c := Color{}
	c.BgMagenta()
	return &c
}
