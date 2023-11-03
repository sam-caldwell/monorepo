package counters

/*
 * LargeCounter.String()
 * (c) 2023 Sam Caldwell.  See license.txt
 *
 * Return the hex string representation of the LargeCounter.
 */

import "encoding/hex"

// String - Return hex string in reverse order
func (c *LargeCounter) String() string {
	return hex.EncodeToString(c.Bytes())
}
