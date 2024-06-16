package ansi

import "fmt"

// Cyan - Set color
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Cyan() *Color {
	if !disabled {
		fmt.Print(CodeFgCyan)
	}
	return c
}

// Cyan - set color and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Cyan() *Color {
	return (&Color{}).Cyan()
}
