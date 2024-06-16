package ansi

import "fmt"

// Indent - indent n spaces
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Indent(n int) *Color {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
	return c
}

// Indent - indent n spaces
//
//	(c) 2023 Sam Caldwell.  MIT License
func Indent(n int) *Color {
	return (&Color{}).Indent(n)
}
