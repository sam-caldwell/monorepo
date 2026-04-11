package logger

import (
	"github.com/sam-caldwell/monorepo/projects/logger/LogLevel"
	"github.com/sam-caldwell/monorepo/projects/ratelimiter/tokensPerSecond"
	"io"
)

// Logger - Top-level logging object
//
//	(c) 2023 Sam Caldwell.  MIT License
type Logger struct {
	useColor  bool
	appName   string
	msgId     string
	level     LogLevel.Value
	ratelimit tokensPerSecond.RateLimiter
	target    io.Writer
}
