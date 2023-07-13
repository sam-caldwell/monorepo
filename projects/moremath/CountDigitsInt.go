package moremath

func CountDigitsInt(n int) int {
	count := 0

	if n == 0 {
		return 1
	}
	//account for - in negative number
	if n < 0 {
		count++
	}
	for n != 0 {
		n /= 10
		count++
	}

	return count
}
