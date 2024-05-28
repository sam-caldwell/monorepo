package ansi

import "fmt"

// Green - Set color
func (c *Color) Green() *Color {
	if !disabled {
		fmt.Print(CodeFgGreen)
	}
	return c
}

// Green - set color and return a new color object
func Green() *Color {
	c := Color{}
	c.Green()
	return &c
}
