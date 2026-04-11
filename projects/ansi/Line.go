package ansi

import (
	"fmt"
	"strings"
)

// Line - Print a line of 'num' 'ch' characters
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Line(ch string, num int) *Color {
	fmt.Println(strings.Repeat(ch, num))
	return c
}

// Line - Print a line of 'num' 'ch' characters
//
//	(c) 2023 Sam Caldwell.  MIT License
func Line(ch string, num int) *Color {
	return (&Color{}).Line(ch, num)
}
