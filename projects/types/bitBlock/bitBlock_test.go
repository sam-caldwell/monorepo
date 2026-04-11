package bitBlock

/*
 * Block Structure Test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Tests for the block struct.
 */

import "testing"

func TestBlockStruct(t *testing.T) {
	block := Block{}
	if block.buffer != nil {
		t.Fatal("expect block buffer to be nil by default")
	}
}
