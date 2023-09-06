package ansi

import "fmt"

// BgRed - Set color
func (c *Color) BgRed() *Color {
	fmt.Print(bgRed)
	return c
}

// BgRed - set color and return a new color object
func BgRed() *Color {
	c := Color{}
	c.BgRed()
	return &c
}
