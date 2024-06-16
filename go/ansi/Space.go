package ansi

import "fmt"

// Space - print a Space character
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Space() *Color {
	if !disabled {
		fmt.Print(Space)
	}
	return c
}
