package tokensPerSecond

import "time"

// NonBlockingCheck - If active, return boolean true if rate limit should allow an operation.
//
//			 If the rate limiter is not active, we will always return true to allow operations.
//			 If the rate limit is exhausted, we will return false, indicating that an operation
//			 should not be allowed.
//
//	         tokensConsumed - represents the number of tokens the operation will consume from the
//	         ratelimit (number of tokens allowed per period).
//
//			 (c) 2023 Sam Caldwell.  MIT License
func (r *RateLimiter) NonBlockingCheck(tokensConsumed int) bool {
	if !r.active {
		return true //Always allow because rate limiter is disabled.
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
	if uint(tokensConsumed) > r.availableTokens {
		return false //We are rate limited.
	}
	r.availableTokens -= uint(tokensConsumed)
	return true
}
