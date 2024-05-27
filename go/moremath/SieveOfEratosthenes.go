package moremath

// SieveOfEratosthenes calculates the number of prime numbers less than n.
func SieveOfEratosthenes(n uint) uint {
	if n <= 2 {
		return 0
	}

	// Initialize the sieve array with true values
	isPrime := make([]bool, n)
	for i := uint(2); i < n; i++ {
		isPrime[i] = true
	}

	// Sieve of Eratosthenes algorithm
	for p := uint(2); p*p < n; p++ {
		if isPrime[p] {
			for multiple := p * p; multiple < n; multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	// Count primes less than n
	count := uint(0)
	for i := uint(2); i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}
