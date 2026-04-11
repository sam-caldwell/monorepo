package convert

import "strconv"

// UnsafeAtoI - ignore any error converting string to integer (0 on error)
//
//	    (c) 2023 Sam Caldwell.  MIT License
func UnsafeAtoI(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return i
}
