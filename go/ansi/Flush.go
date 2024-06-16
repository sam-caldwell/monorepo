package ansi

import (
    "os"
)

// Flush - Flush StdOut buffer
//
//	(c) 2023 Sam Caldwell.  MIT License
func (c *Color) Flush() *Color {
    defer func() { _ = os.Stdout.Sync() }()
    return c
}

// Flush - Flush StdOut buffer
//
//	(c) 2023 Sam Caldwell.  MIT License
func Flush() *Color {
    return (&Color{}).Flush()
}
