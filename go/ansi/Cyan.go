package ansi

import "fmt"

// Cyan - Set color
func (c *Color) Cyan() *Color {
	if !c.disabled {
		fmt.Print(fgCyan)
	}
	return c
}

// Cyan - set color and return a new color object
func Cyan() *Color {
	c := Color{}
	c.Cyan()
	return &c
}
