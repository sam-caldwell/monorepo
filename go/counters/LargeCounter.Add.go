package counters

/*
 * LargeCounter.Add() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * A LargeCounter method to add a uint64 to a given counter.
 */

import "math"

func (c *LargeCounter) Add(value uint64) (err error) {
	return c.carryValue(0, value)
}

func (c *LargeCounter) carryValue(pos uint, v uint64) (err error) {
	if pos < uint(len(c.v)) {
		n := math.MaxUint64 - v
		if n > 0 {
			c.v[pos] = 0
			return c.carryValue(pos+1, n)
		}
	}
	return nil
}
