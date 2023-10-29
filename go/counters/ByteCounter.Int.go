package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt

 * ByteCounter will create an array of n bytes
 * and allow the count to be incremented 0...255
 * for each byte, carrying 1 to the n+1 byte.
 */

import (
	"math/big"
)

// Int - return the numeric value of our counter state (using big int)
func (c *ByteCounter) Int() *big.Int {
	c.reverse()
	i := big.Int{}
	i.SetBytes(c.v)
	return &i
}
