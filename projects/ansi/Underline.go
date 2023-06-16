package ansi

import "fmt"

// Underline - Set format attribute
func (c *Color) Underline() *Color {
	fmt.Print(underline)
	return c
}

// Underline - set format attribute and return a new color object
func Underline() *Color {
	c := Color{}
	c.Underline()
	return &c
}
