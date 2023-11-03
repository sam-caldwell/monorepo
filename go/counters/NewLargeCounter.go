package counters

/*
 * NewLargeCounter
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A function to create an initialized LargeCounter
 */

import "fmt"

// NewLargeCounter - Create a new LargeCounter object if >=64 bits with a size (in bits)
func NewLargeCounter(sz int) (*LargeCounter, error) {
	if sz <= 64 {
		return nil, fmt.Errorf("bounds check error")
	}
	o := make(LargeCounter, sz)
	return &o, nil
}
