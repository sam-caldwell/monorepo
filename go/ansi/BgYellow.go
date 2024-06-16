package ansi

import "fmt"

// BgYellow - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) BgYellow() *Color {
	if !disabled {
		fmt.Print(CodeBgYellow)
	}
	return c
}

// BgYellow - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func BgYellow() *Color {
	return (&Color{}).BgYellow()
}
