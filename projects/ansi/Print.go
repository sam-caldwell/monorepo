package ansi

import "fmt"

// Print - print text to stdout
func (c *Color) Print(message string) *Color {
	fmt.Print(message)
	return c
}

// Print - print text to stdout and return color object
func Print(message string) *Color {
	c := Color{}
	c.Print(message)
	return &c
}
