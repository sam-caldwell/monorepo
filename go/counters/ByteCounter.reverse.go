package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt

 * ByteCounter will create an array of n bytes
 * and allow the count to be incremented 0...255
 * for each byte, carrying 1 to the n+1 byte.
 */

func (c *ByteCounter) reverse() {
	sz := len(c.v)
	rhs := sz
	sz /= 2
	for lhs := 0; lhs < sz; lhs++ {
		rhs--
		c.v[lhs] = c.v[rhs]
	}
}
