package ansi

import "fmt"

// Clear - Set format attribute
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Clear() *Color {
	if !disabled {
		fmt.Print(CodeClear)
	}
	return c
}

// Clear - set format attribute and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Clear() *Color {
	return (&Color{}).Clear()
}
