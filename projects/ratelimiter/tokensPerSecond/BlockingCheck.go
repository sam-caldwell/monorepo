package tokensPerSecond

import "time"

// BlockingCheck - If active, determine if rate limit should block an operation until enough credits exist.
//
//	          If the rate limiter is not active, we will always return true to allow operations.
//
//		      tokensConsumed - represents the number of tokens the operation will consume from the
//		      ratelimit (number of tokens allowed per period).
//
//			(c) 2023 Sam Caldwell.  MIT License
func (r *RateLimiter) BlockingCheck(tokensConsumed int) {
	if !r.active {
		return //Always allow because rate limiter is disabled.
	}
	now := time.Now()
	elapsedTime := now.Sub(r.lastAccessTime).Seconds()

	tokenCredits := uint(elapsedTime) * r.rateLimit
	if tokenCredits > r.rateLimit {
		tokenCredits = r.rateLimit
	}
	r.availableTokens += tokenCredits
	if r.availableTokens > r.rateLimit {
		r.availableTokens = r.rateLimit
	}

	for tokensNeeded := r.availableTokens - uint(tokensConsumed); tokensNeeded > 0; tokensNeeded -= r.availableTokens {
		time.Sleep(time.Second * 1)
		now = time.Now()
		elapsedTime = now.Sub(r.lastAccessTime).Seconds()
		tokenCredits = uint(elapsedTime) * r.rateLimit
		if tokenCredits > r.rateLimit {
			tokenCredits = r.rateLimit
		}
		r.availableTokens += tokenCredits
		if r.availableTokens > r.rateLimit {
			r.availableTokens = r.rateLimit
		}
	}
}
