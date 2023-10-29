package counters

import (
	"fmt"
	"math/big"
)

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt

 * ByteCounter will create an array of n bytes
 * and allow the count to be incremented 0...255
 * for each byte, carrying 1 to the n+1 byte.
 */

type ByteCounter struct {
	v []byte // byte array used for the counter.
}

func NewByteCounter(sz int) (ByteCounter, error) {
	var b ByteCounter
	if sz <= 0 {
		return b, fmt.Errorf("ByteCounter size must be > 0")
	}
	b.v = make([]byte, sz)
	return b, nil
}

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
		err = fmt.Errorf("overflow error")
	}
	return err
}

// Increment - increment the byte counter (carry 1s as needed)
func (c *ByteCounter) Increment() error {
	return c.inc(0)
}

func (c *ByteCounter) reverseBytes() {
	fmt.Println("reverseBytes")
	o := make([]byte, len(c.v))
	copy(c.v, o)
	for i := 0; i < len(o); i++ {
		c.v[i] = o[(len(o)-i)-1]
	}
}

// Bytes - return the byte string value of our counter state.
func (c *ByteCounter) Bytes() []byte {
	o := make([]byte, len(c.v))
	for i := 0; i < len(o); i++ {
		o[i] = c.v[(len(o)-i)-1]
	}
	return c.v
}

// String - return the string value of our counter state.
func (c *ByteCounter) String() string {
	c.reverseBytes()
	return (string)(c.v)
}

// Int - return the numeric value of our counter state (using big int)
func (c *ByteCounter) Int() *big.Int {
	c.reverseBytes()
	i := big.Int{}
	i.SetBytes(c.v)
	return &i
}
