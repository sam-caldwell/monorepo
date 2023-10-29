package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt

 * ByteCounter will create an array of n bytes
 * and allow the count to be incremented 0...255
 * for each byte, carrying 1 to the n+1 byte.
 */

func (c *ByteCounter) reverse() {
	o := make([]byte, len(c.v))
	copy(c.v, o)
	for i := 0; i < len(o); i++ {
		c.v[i] = o[(len(o)-i)-1]
	}
}
