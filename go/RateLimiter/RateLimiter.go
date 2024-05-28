package ratelimiter

import "time"

// RateLimiter - Mechanism for rate limiting
//
//	(c) 2023 Sam Caldwell.  MIT License
type RateLimiter struct {
	active          bool
	rateLimit       uint //How many tokens per second will we allow?
	lastAccessTime  time.Time
	availableTokens uint
}
