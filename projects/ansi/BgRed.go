package ansi

import "fmt"

// BgRed - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) BgRed() *Color {
	if !disabled {
		fmt.Print(CodeBgRed)
	}
	return c
}

// BgRed - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func BgRed() *Color {
	return (&Color{}).BgRed()
}
