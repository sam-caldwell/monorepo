package ansi

import "fmt"

// Blink - Set format attribute
func (c *Color) Blink() *Color {
	fmt.Print(blink)
	return c
}

// Blink - set format attribute and return a new color object
func Blink() *Color {
	c := Color{}
	c.Blink()
	return &c
}
