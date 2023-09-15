package ansi

import "fmt"

// Red - Set color
func (c *Color) Red() *Color {
	if !c.disabled {
		fmt.Print(fgRed)
	}
	return c
}

// Red - set color and return a new color object
func Red() *Color {
	c := Color{}
	c.Red()
	return &c
}
