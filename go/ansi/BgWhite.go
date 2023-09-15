package ansi

import "fmt"

// BgWhite - Set color
func (c *Color) BgWhite() *Color {
	if !disabled {
		fmt.Print(bgWhite)
	}
	return c
}

// BgWhite - set color and return a new color object
func BgWhite() *Color {
	c := Color{}
	c.BgWhite()
	return &c
}
