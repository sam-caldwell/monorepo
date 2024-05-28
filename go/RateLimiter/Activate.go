package ratelimiter

// Activate - Enable rate limiting
//
//	(c) 2023 Sam Caldwell.  MIT License
func (r *RateLimiter) Activate() {
	r.active = true
}
