package ansi

import "fmt"

// Yellow - set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Yellow() *Color {
	if !disabled {
		fmt.Print(CodeFgYellow)
	}
	return c
}

// Yellow - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Yellow() *Color {
	return (&Color{}).Yellow()
}
