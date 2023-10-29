package counters

/*
 * ByteCounter
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * ByteCounter tests
 */

import (
	"testing"
)

func TestByteCounter_Struct(t *testing.T) {
	var b ByteCounter
	if b.v != nil {
		t.Fatal("b.v should be nil")
	}
}
