package ansi

import "fmt"

// Space - print a Space character
func (c *Color) Space() *Color {
	fmt.Print(Space)
	return c
}
