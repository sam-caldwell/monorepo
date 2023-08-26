package ansi

import "fmt"

// Up - move cursor n units
func (c *Color) Up(n int) *Color {
	fmt.Printf(moveUp, n)
	return c
}

// Up - move cursor n units and return a new color object
func Up(n int) *Color {
	c := Color{}
	c.Up(n)
	return &c
}
