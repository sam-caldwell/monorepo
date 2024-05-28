package logger

import (
	"io"
)

// Logger - Top-level logging object
//
//	(c) 2023 Sam Caldwell.  MIT License
type Logger struct {
	target    io.Writer
	rateLimit uint
}
