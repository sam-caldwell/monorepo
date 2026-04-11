package counters

import (
	"bytes"
	"testing"
)

/*
 * ByteCounter.reverse()
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Reverse the order of bytes
 */

func TestByteCounter_reverse(t *testing.T) {
	c := ByteCounter{
		v: []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09},
	}
	c.reverse()
	expected := []byte{0x09, 0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0x00}
	if bytes.Equal(c.v, expected) {
		t.Fatal("reversed string is wrong")
	}
}
