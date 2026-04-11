package ansi

import (
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02-15:04.0500 "
)

// Time - print current time
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Time() *Color {
	fmt.Print(time.Now().Format(timeFormat))
	return c
}

// Time - print current time
//
//	(c) 2023 Sam Caldwell.  MIT License
func Time() *Color {
    return (&Color{}).Time()
}
