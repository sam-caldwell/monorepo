package tokensPerSecond

// GetLimit - Return the current rate limit (number allowed tokens per second)
//
//	(c) 2023 Sam Caldwell.  MIT License
func (r *RateLimiter) GetLimit() uint {
	return r.rateLimit
}
