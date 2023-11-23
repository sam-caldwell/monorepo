package counters

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * ByteCounter Adder Function
 */

func (c *ByteCounter) Add(value int) (err error) {

	if value < 0 {
		return nil
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	remainder := value % 256

	// Add the remainder to the counter
	for pos := 0; pos < len(c.v); pos++ {
		newValue := int(c.v[pos]) + remainder
		c.v[pos] = byte(newValue % 256)
		remainder = newValue / 256
		if remainder <= 0 {
			return nil
		}
	}
	return fmt.Errorf(errors.OverflowError)
}
