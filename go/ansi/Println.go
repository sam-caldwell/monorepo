package ansi

import "fmt"

// Println - print text to stdout
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Println(message string) *Color {
	if !disabled {
		fmt.Println(message)
	}
	return c
}

// Println - print text to stdout and return color object
//
//	(c) 2023 Sam Caldwell.  MIT License
func Println(message string) *Color {
	return (&Color{}).Println(message)
}
