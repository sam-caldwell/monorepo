package ansi

import "fmt"

// Dim - decrease intensity
func (c *Color) Dim() *Color {
	if !disabled {
		fmt.Print(CodeDim)
	}
	return c
}

// Dim - decrease intensity and return a new color object
func Dim() *Color {
	c := Color{}
	c.Dim()
	return &c
}
