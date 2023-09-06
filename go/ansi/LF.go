package ansi

import "fmt"

// LF - print a line feed char
func (c *Color) LF() *Color {
	fmt.Print(LineFeed)
	return c
}

// LF - print a line feed char
func LF() *Color {
	c := Color{}
	c.LF()
	return &c
}
