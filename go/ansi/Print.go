package ansi

import "fmt"

// Print - print text to stdout
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Print(message string) *Color {
	if !disabled {
		fmt.Print(message)
	}
	return c
}

// Print - print text to stdout and return color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Print(message string) *Color {
    return (&Color{}).Print(message)
}
