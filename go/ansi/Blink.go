package ansi

import "fmt"

// Blink - Set format attribute
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Blink() *Color {
	if !disabled {
		fmt.Print(CodeBlink)
	}
	return c
}

// Blink - set format attribute and return a new color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Blink() *Color {
	return (&Color{}).Blink()
}
