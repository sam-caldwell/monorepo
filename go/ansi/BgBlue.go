package ansi

import "fmt"

// BgBlue - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) BgBlue() *Color {
	if !disabled {
		fmt.Print(CodeBgBlue)
	}
	return c
}

// BgBlue - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func BgBlue() *Color {
	return (&Color{}).BgBlue()
}
