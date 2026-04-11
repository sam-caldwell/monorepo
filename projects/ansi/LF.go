package ansi

import "fmt"

// LF - print a line feed char
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) LF() *Color {
	fmt.Print(LineFeed)
	return c
}

// LF - print a line feed char
//
//	(c) 2023 Sam Caldwell.  MIT License
func LF() *Color {
	return (&Color{}).LF()
}
