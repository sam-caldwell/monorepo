package ratelimiter

// GetRemaining - Return the available tokens
//
//	(c) 2023 Sam Caldwell.  MIT License
func (r *RateLimiter) GetRemaining() uint {
    return r.availableTokens
}
