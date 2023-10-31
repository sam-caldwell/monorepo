package counters

import "fmt"

/*
 * ByteCounter.Set
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Set() method will allow a given byte position
 * to be set to a given value
 */

func (c *ByteCounter) Set(pos int, value byte) (err error) {
	if c.v == nil {
		c.v = make([]byte, pos+1)
	}
	if (pos >= len(c.v)) || (pos < 0) {
		err = fmt.Errorf("index error")
	} else {
		c.v[pos] = value
	}
	return
}
