package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"io"
)

// Logger - Top-level logging object
//
//	(c) 2023 Sam Caldwell.  MIT License
type Logger struct {
	appName   string
	msgId     string
	level     LogLevel.Value
	rateLimit uint
	target    io.Writer
}
