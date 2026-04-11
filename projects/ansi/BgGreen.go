package ansi

import "fmt"

// BgGreen - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) BgGreen() *Color {
	if !disabled {
		fmt.Print(CodeBgGreen)
	}
	return c
}

// BgGreen - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func BgGreen() *Color {
    return (&Color{}).BgGreen()
}
