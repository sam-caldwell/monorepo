package counters

/*
 * LargeCounter test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test the basic type.
 */

import "testing"

func TestLargeCounter(t *testing.T) {
	var o LargeCounter
	if o != nil {
		t.Fatal("expect nil initial state")
	}
}
