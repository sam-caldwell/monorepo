package ansi

import "fmt"

// BgWhite - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) BgWhite() *Color {
	if !disabled {
		fmt.Print(CodeBgWhite)
	}
	return c
}

// BgWhite - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func BgWhite() *Color {
	return (&Color{}).BgWhite()
}
