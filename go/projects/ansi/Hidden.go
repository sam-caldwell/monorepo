package ansi

import "fmt"

// Hidden - Set format attribute
func (c *Color) Hidden() *Color {
	fmt.Print(hiddenText)
	return c
}

// Hidden - set format attribute and return a new color object
func Hidden() *Color {
	c := Color{}
	c.Hidden()
	return &c
}
