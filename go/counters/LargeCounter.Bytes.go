package counters

/*
 * LargeCounter.Bytes() method
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Return the []byte string representation of the LargeCounter.
 */
import "encoding/binary"

// Bytes - Return the bytes in reverse order
func (c *LargeCounter) Bytes() []byte {
	out := make([]byte, 8*len(c.v))
	for i := len(c.v) - 1; i >= 0; i-- {
		this := make([]byte, 8)
		binary.LittleEndian.PutUint64(this, c.v[i])
		out = append(out, this...)
	}
	return out
}
