package ansi

import "fmt"

// BgGreen - Set color
func (c *Color) BgGreen() *Color {
	fmt.Print(bgGreen)
	return c
}

// BgGreen - set color and return a new color object
func BgGreen() *Color {
	c := Color{}
	c.BgGreen()
	return &c
}
