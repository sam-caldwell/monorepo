package ansi

import "fmt"

// Tab - print a Tab char
func (c *Color) Tab() *Color {
	fmt.Print(Tab)
	return c
}
