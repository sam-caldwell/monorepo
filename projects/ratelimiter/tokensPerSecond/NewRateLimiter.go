package tokensPerSecond

import "time"

// NewRateLimiter - Create, configure and return a new rate limiter object
//
//	(c) 2023 Sam Caldwell.  MIT License
func NewRateLimiter(limit int) RateLimiter {
	return RateLimiter{
		rateLimit:       uint(limit),
		lastAccessTime:  time.Now(),
		availableTokens: uint(limit),
		active:          true,
	}
}
