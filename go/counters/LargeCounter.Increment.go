package counters

/*
 * LargeCounter.Increment()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Increment the large counter value, carrying
 * any overflow to the next element in the array.
 */

import (
	"math"
)

// Increment - increment the large counter by one.
func (c *LargeCounter) Increment() {
	for i := 0; i < len(*c); i++ {
		if (*c)[i] == math.MaxUint64 {
			(*c)[i] = 0
		} else {
			(*c)[i]++
			return
		}
	}
	return
}
