package ansi

import "fmt"

// Tab - print a tab char
func (c *Color) Tab() *Color {
	fmt.Print(tab)
	return c
}
