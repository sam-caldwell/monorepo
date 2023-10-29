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
	o := make([]byte, len(c.v))
	for i := 0; i < len(o); i++ {
		o[i] = c.v[(len(o)-i)-1]
	}
	return c.v
}
