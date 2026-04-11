package ansi

import "fmt"

// Tab - print a Tab char
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Tab() *Color {
	fmt.Print(Tab)
	return c
}
