package tokensPerSecond

// Disable - disable the rate limiter
//
//	 This will result in all activity being allowed because no rate limiting
//	 will be enforced.
//
//		(c) 2023 Sam Caldwell.  MIT License
func (r *RateLimiter) Disable() {
	r.active = false
}
