package ansi

import "fmt"

// Blink - Set format attribute
func (c *Color) Blink() *Color {
	if !disabled {
		fmt.Print(CodeBlink)
	}
	return c
}

// Blink - set format attribute and return a new color object
func Blink() *Color {
	c := Color{}
	c.Blink()
	return &c
}
