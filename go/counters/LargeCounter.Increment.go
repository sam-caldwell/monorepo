package counters

/*
 * LargeCounter.Increment()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Increment the large counter value, carrying
 * any overflow to the next element in the array.
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"math"
)

// Increment - increment the large counter by one.
func (c *LargeCounter) Increment() error {
	return c.carry(0)
}

// carry - recursively carry the value at element pos if overflow or increment.
func (c *LargeCounter) carry(pos uint) error {
	if pos > uint(len(c.v)) {
		return fmt.Errorf(errors.OverflowError)
	}
	if c.v[pos] == math.MaxUint64 {
		c.v[pos] = 0
		return c.carry(pos + 1)
	}
	c.v[pos]++
	return nil
}
