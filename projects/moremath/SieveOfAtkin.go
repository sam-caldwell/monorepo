package moremath

import (
	"math"
)

// SieveOfAtkin calculates the number of prime numbers less than n.
func SieveOfAtkin(n uint) uint {
	if n <= 2 {
		return 0
	}

	// Initialize the sieve array with false values
	isPrime := make([]bool, n+1)

	// Sieve of Atkin algorithm
	sqrt := uint(math.Sqrt(float64(n)))
	for x := uint(1); x <= sqrt; x++ {
		for y := uint(1); y <= sqrt; y++ {
			n1 := 4*x*x + y*y
			if n1 <= n && (n1%12 == 1 || n1%12 == 5) {
				isPrime[n1] = !isPrime[n1]
			}

			n2 := 3*x*x + y*y
			if n2 <= n && n2%12 == 7 {
				isPrime[n2] = !isPrime[n2]
			}

			n3 := 3*x*x - y*y
			if x > y && n3 <= n && n3%12 == 11 {
				isPrime[n3] = !isPrime[n3]
			}
		}
	}

	// Mark all multiples of squares of primes as false
	for r := uint(5); r <= sqrt; r++ {
		if isPrime[r] {
			for k := r * r; k <= n; k += r * r {
				isPrime[k] = false
			}
		}
	}

	// Count primes less than n
	count := uint(2) // 2 and 3 are primes
	if n > 2 {
		count = 1 // 2 is a prime
	}
	if n > 3 {
		count++ // 3 is a prime
	}

	for i := uint(5); i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}
