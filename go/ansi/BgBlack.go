package ansi

import "fmt"

// BgBlack - Set color
func (c *Color) BgBlack() *Color {
	if !disabled {
		fmt.Print(CodeBgBlack)
	}
	return c
}

// BgBlack - set color and return a new color object
func BgBlack() *Color {
	c := Color{}
	c.BgBlack()
	return &c
}
