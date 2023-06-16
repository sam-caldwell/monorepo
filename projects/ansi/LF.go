package ansi

import "fmt"

// LF - print a line feed char
func (c *Color) LF() *Color {
	fmt.Print(lineFeed)
	return c
}
