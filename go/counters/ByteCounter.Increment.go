package counters

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

//
//// Increment - increment the byte counter (carry 1s as needed)
//func (c *ByteCounter) Increment() error {
//	return c.inc(0)
//}

// Increment - increment the large counter by one.
func (c *ByteCounter) Increment() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	for i := 0; i < len(c.v); i++ {
		c.v[i]++
		if c.v[i] != 0 {
			return nil
		}
	}
	return fmt.Errorf(errors.OverflowError)
}

// FastIncrement - increment the large counter by one (without locking)
func (c *ByteCounter) FastIncrement() error {
	for i := 0; i < len(c.v); i++ {
		c.v[i]++
		if c.v[i] != 0 {
			return nil
		}
	}
	return fmt.Errorf(errors.OverflowError)
}
