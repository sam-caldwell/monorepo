package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"io"
	"time"
)

// Logger - Top-level logging object
//
//	(c) 2023 Sam Caldwell.  MIT License
type Logger struct {
	useColor bool
	appName  string
	msgId    string
	level    LogLevel.Value
	target   io.Writer
}

// RateLimiter - Mechanism for rate limiting
type RateLimiter struct {
	rateLimit      uint //How many tokens per second will we allow?
	lastAccessTime time.Time
	availableBytes uint
}

func (r *RateLimiter) SetRateLimit(quota uint) {
	r.rateLimit = quota
}

func (r *RateLimiter) Reset() {
	r.availableBytes = r.rateLimit
	r.lastAccessTime = time.Now()
}
