package ansi

import "fmt"

// Green - Set color
func (c *Color) Green() *Color {
	fmt.Print(fgGreen)
	return c
}

// Green - set color and return a new color object
func Green() *Color {
	c := Color{}
	c.Green()
	return &c
}
