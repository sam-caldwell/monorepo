package ansi

import "fmt"

// BgMagenta - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) BgMagenta() *Color {
	if !disabled {
		fmt.Print(CodeBgMagenta)
	}
	return c
}

// BgMagenta - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func BgMagenta() *Color {
    return (&Color{}).BgMagenta()
}
