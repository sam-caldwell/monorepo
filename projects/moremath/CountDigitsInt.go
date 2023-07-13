package moremath

func CountDigitsInt(n int) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	if n < 0 {
		count++
	}

	return count
}
