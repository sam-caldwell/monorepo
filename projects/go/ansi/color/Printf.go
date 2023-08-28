package ansi

import "fmt"

// Printf - print text to stdout
func (c *Color) Printf(format string, a ...any) *Color {
	fmt.Printf(format, a...)
	return c
}

// Printf - print text to stdout and return color object
func Printf(format string, a ...any) *Color {
	c := Color{}
	c.Printf(format, a...)
	return &c
}
