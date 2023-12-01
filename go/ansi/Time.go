package ansi

import (
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02-15:04.0500 "
)

// Time - print current time
func (c *Color) Time() *Color {
	fmt.Print(time.Now().Format(timeFormat))
	return c
}

// Time - print current time
func Time() *Color {
	c := Color{}
	c.Time()
	return &c
}
