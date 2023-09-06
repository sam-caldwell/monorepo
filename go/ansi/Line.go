package ansi

import (
	"fmt"
	"strings"
)

// Line - Print a line of 'num' 'ch' characters
func (c *Color) Line(ch string, num int) *Color {
	fmt.Println(strings.Repeat(ch, num))
	return c
}

// Line - Print a line of 'num' 'ch' characters
func Line(ch string, num int) *Color {
	c := Color{}
	c.Line(ch, num)
	return &c
}
