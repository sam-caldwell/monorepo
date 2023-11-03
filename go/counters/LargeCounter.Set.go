package counters

/*
 * LargeCounter.Set() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Set a specific element in the counter (initialize if not initialized)
 */

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Set - Set element at pos to value and initialize the c.v if nil
func (c *LargeCounter) Set(pos int, value uint64) (err error) {
	if *c == nil {
		*c = make([]uint64, pos+1)
	}
	if pos >= len(*c) {
		return fmt.Errorf(errors.OverflowError)
	}
	(*c)[pos] = value
	return err
}
