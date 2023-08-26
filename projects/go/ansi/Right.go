package ansi

import "fmt"

// Right - move cursor n units
func (c *Color) Right(n int) *Color {
	fmt.Printf(moveRight, n)
	return c
}

// Right - move cursor n units and return a new color object
func Right(n int) *Color {
	c := Color{}
	c.Right(n)
	return &c
}
