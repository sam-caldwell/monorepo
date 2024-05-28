package tokensPerSecond

// IsActive - return active state of the rate limiter.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (r *RateLimiter) IsActive() bool {
	return r.active
}
