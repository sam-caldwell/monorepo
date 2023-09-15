package ansi

import "fmt"

// Bold - Set format attribute
func (c *Color) Bold() *Color {
	if !disabled {
		fmt.Print(bold)
	}
	return c
}

// Bold - set format attribute and return a new color object
func Bold() *Color {
	c := Color{}
	c.Bold()
	return &c
}
