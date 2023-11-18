package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt

 * ByteCounter will create an array of n bytes
 * and allow the count to be incremented 0...255
 * for each byte, carrying 1 to the n+1 byte.
 */

// Bytes - return the byte string value of our counter state.
func (c *ByteCounter) Bytes() []byte {
	c.lock.Lock()
	defer c.lock.Unlock()
	return reverseBytes(c.v)
}

func (c *ByteCounter) RawBytes() []byte {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.v
}
