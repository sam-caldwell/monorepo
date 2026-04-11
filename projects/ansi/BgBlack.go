package ansi

import "fmt"

// BgBlack - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) BgBlack() *Color {
	if !disabled {
		fmt.Print(CodeBgBlack)
	}
	return c
}

// BgBlack - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func BgBlack() *Color {
	return (&Color{}).BgBlack()
}
