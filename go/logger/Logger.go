package logger

import (
	ratelimiter "github.com/sam-caldwell/monorepo/go/RateLimiter"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
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
	ratelimit ratelimiter.RateLimiter
	target    io.Writer
}
