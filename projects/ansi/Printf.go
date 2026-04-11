package ansi

import "fmt"

// Printf - print text to stdout
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Printf(format string, a ...any) *Color {
	if !disabled {
		fmt.Printf(format, a...)
	}
	return c
}

// Printf - print text to stdout and return color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Printf(format string, a ...any) *Color {
	return (&Color{}).Printf(format, a...)
}
