package ansi

import "fmt"

// BgCyan - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) BgCyan() *Color {
	if !disabled {
		fmt.Print(CodeBgCyan)
	}
	return c
}

// BgCyan - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func BgCyan() *Color {
    return (&Color{}).BgCyan()
}
