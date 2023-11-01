package counters

import (
	"encoding/hex"
	"strings"
)

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt

 * ByteCounter will create an array of n bytes
 * and allow the count to be incremented 0...255
 * for each byte, carrying 1 to the n+1 byte.
 */

// String - return the string value of our counter state.
func (c *ByteCounter) String() string {
	//c.reverse()
	return strings.ToUpper(hex.EncodeToString(c.Bytes()))
}
