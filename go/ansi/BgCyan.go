package ansi

import "fmt"

// BgCyan - Set color
func (c *Color) BgCyan() *Color {
	if !disabled {
		fmt.Print(bgCyan)
	}
	return c
}

// BgCyan - set color and return a new color object
func BgCyan() *Color {
	c := Color{}
	c.BgCyan()
	return &c
}
