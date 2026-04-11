package ansi

import "fmt"

// Dim - decrease intensity
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Dim() *Color {
	if !disabled {
		fmt.Print(CodeDim)
	}
	return c
}

// Dim - decrease intensity and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Dim() *Color {
	c := Color{}
	c.Dim()
	return &c
}
