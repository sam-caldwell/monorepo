package misc

/*
 * projects/misc/PickSmallestInt.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Given one or more integers, this function will
 * return the smallest.
 */

// PickSmallestInt - return the smallest integer
func PickSmallestInt(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	smallest := numbers[0]
	for _, num := range numbers {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}
