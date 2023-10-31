package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt

 * ByteCounter will create an array of n bytes
 * and allow the count to be incremented 0...255
 * for each byte, carrying 1 to the n+1 byte.
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

func (c *ByteCounter) inc(n int) error {
	if c.v[n] == 255 {
		c.v[n] = 0
		return c.carry(n + 1)
	} else {
		c.v[n]++
	}
	return nil
}

func (c *ByteCounter) carry(n int) (err error) {
	if n < len(c.v) {
		return c.inc(n)
	} else {
		err = fmt.Errorf(errors.OverflowError)
	}
	return err
}

// Increment - increment the byte counter (carry 1s as needed)
func (c *ByteCounter) Increment() error {
	return c.inc(0)
}
