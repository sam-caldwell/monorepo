package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt

 * ByteCounter will create an array of n bytes
 * and allow the count to be incremented 0...255
 * for each byte, carrying 1 to the n+1 byte.
 */

import (
	"fmt"
)

func NewByteCounter(sz int) (ByteCounter, error) {
	var b ByteCounter
	if sz <= 0 {
		return b, fmt.Errorf("ByteCounter size must be > 0")
	}
	b.v = make([]byte, sz)
	return b, nil
}
