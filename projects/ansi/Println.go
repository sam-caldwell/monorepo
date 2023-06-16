package ansi

import "fmt"

// Println - print text to stdout
func (c *Color) Println(message string) *Color {
	fmt.Println(message)
	return c
}

// Println - print text to stdout and return color object
func Println(message string) *Color {
	c := Color{}
	c.Println(message)
	return &c
}
