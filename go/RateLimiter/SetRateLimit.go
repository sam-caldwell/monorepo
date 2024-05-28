package ratelimiter

// SetRateLimit - Set the number of tokens allowed per second
//
//	(c) 2023 Sam Caldwell.  MIT License
func (r *RateLimiter) SetRateLimit(quota uint) {
	r.rateLimit = quota
}
