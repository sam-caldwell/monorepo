package tokensPerSecond

import "time"

// Reset - Reset the rate limiter
//
//	(c) 2023 Sam Caldwell.  MIT License
func (r *RateLimiter) Reset() {
	r.availableTokens = r.rateLimit
	r.lastAccessTime = time.Now()
}
