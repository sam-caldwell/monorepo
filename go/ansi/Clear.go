package ansi

import "fmt"

// Clear - Set format attribute
func (c *Color) Clear() *Color {
	if !disabled {
		fmt.Print(clear)
	}
	return c
}

// Clear - set format attribute and return a new color object
func Clear() *Color {
	c := Color{}
	c.Clear()
	return &c
}
