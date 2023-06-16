package ansi

import "fmt"

// BgYellow - Set color
func (c *Color) BgYellow() *Color {
	fmt.Print(bgYellow)
	return c
}

// BgYellow - set color and return a new color object
func BgYellow() *Color {
	c := Color{}
	c.BgYellow()
	return &c
}
