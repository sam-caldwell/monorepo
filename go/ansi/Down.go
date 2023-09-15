package ansi

import "fmt"

// Down - move cursor n units
func (c *Color) Down(n int) *Color {
	if !disabled {
		fmt.Printf(moveDown, n)
	}
	return c
}

// Down - move cursor n units and return a new color object
func Down(n int) *Color {
	c := Color{}
	c.Down(n)
	return &c
}
