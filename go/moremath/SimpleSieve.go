package moremath

// SimpleSieve finds all prime numbers up to a given limit
func SimpleSieve(limit uint) []uint {
	isPrime := make([]bool, limit+1)
	for i := uint(2); i <= limit; i++ {
		isPrime[i] = true
	}

	for p := uint(2); p*p <= limit; p++ {
		if isPrime[p] {
			for multiple := p * p; multiple <= limit; multiple += p {
				isPrime[multiple] = false
			}
		}
	}

	var primes []uint
	for p := uint(2); p <= limit; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}

	return primes
}
