package moremath

import "math"

// SegmentedSieve counts the number of primes less than n using the Segmented Sieve of Eratosthenes
func SegmentedSieve(n uint) uint {
	if n <= 2 {
		return 0
	}

	limit := uint(math.Sqrt(float64(n))) + 1
	primes := SimpleSieve(limit)

	count := uint(0)
	low := limit
	high := 2 * limit

	for low < n {
		if high > n {
			high = n
		}

		isPrime := make([]bool, high-low)
		for i := range isPrime {
			isPrime[i] = true
		}

		for _, p := range primes {
			lowLimit := (low + p - 1) / p * p
			if lowLimit < p*p {
				lowLimit = p * p
			}
			for j := lowLimit; j < high; j += p {
				isPrime[j-low] = false
			}
		}

		for i := low; i < high; i++ {
			if isPrime[i-low] {
				count++
			}
		}

		low = high
		high += limit
	}

	return count + uint(len(primes))
}
