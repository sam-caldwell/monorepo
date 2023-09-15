package ansi

import "fmt"

// BgBlue - Set color
func (c *Color) BgBlue() *Color {
	if !c.disabled {
		fmt.Print(bgBlue)
	}
	return c
}

// BgBlue - set color and return a new color object
func BgBlue() *Color {
	c := Color{}
	c.BgBlue()
	return &c
}
