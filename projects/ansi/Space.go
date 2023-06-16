package ansi

import "fmt"

// Space - print a space character
func (c *Color) Space() *Color {
	fmt.Print(space)
	return c
}
