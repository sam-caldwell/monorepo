package ansi

import "fmt"

// Bold - Set format attribute
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Bold() *Color {
	if !disabled {
		fmt.Print(CodeBold)
	}
	return c
}

// Bold - set format attribute and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Bold() *Color {
	return (&Color{}).Bold()
}
