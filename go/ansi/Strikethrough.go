package ansi

import "fmt"

// Strikethrough - Set format attribute
func (c *Color) Strikethrough() *Color {
	if !disabled {
		fmt.Print(strikethrough)
	}
	return c
}

// Strikethrough - set format attribute and return a new color object
func Strikethrough() *Color {
	c := Color{}
	c.Strikethrough()
	return &c
}
