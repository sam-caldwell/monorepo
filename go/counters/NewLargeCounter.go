package counters

/*
 * NewLargeCounter
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A function to create an initialized LargeCounter
 */

import (
	"fmt"
	"math"
)

// NewLargeCounter - Create a new LargeCounter object if >=64 bits with a size (in bits)
func NewLargeCounter(sz int) (*LargeCounter, error) {
	if sz < 64 {
		return nil, fmt.Errorf("bounds check error")
	}
	o := make(LargeCounter, uint64(math.Ceil(float64(sz)/float64(64))))
	return &o, nil
}
