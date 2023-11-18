package counters

import "bytes"

// Equal - Compare two ByteCounters and return true if equal
func (c *ByteCounter) Equal(rhs *ByteCounter) bool {

	c.lock.Lock()
	rhs.lock.Lock()

	defer c.lock.Unlock()
	defer rhs.lock.Unlock()

	return bytes.Equal(c.v, rhs.v)
}
