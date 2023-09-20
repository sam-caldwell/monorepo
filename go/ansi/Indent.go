package ansi

import "fmt"

// Indent - indent n spaces
func (c *Color) Indent(n int) *Color {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
	return c
}

// Indent - indent n spaces
func Indent(n int) *Color {
	c := Color{}
	c.Indent(n)
	return &c
}
