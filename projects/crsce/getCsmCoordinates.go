package crsce

/*
 * projects/crypto/crsce/getCsmCoordinates.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file provides a function that will take a
 * linear address of a CSM element and the size of CSM
 * and return the appropriate (x,y) coordinate pair.
 */

// getCsmCoordinates - Given index i and CSM size n, return the (x,y) coordinate pair.
func getCsmCoordinates(i, n int) (x int, y int) {
	y = i / n
	x = i - y*n
	return x, y
}
