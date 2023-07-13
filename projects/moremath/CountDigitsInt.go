package moremath

/*
 * projects/moremath/CountDigitsInt.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides a CountDigitsInt() function
 * which will return the number of digits in a given
 * number (n).
 *
 * See README.md
 */

// CountDigitsInt - return number of digits in integer (n)
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
