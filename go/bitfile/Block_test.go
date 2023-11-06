package bitfile

/*
 * Block Structure Test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Tests for the block struct.
 */

import "testing"

func TestBlockStruct(t *testing.T) {
	var b Block
	if b.buffer != nil {
		t.Fatal("expect block to have nil buffer initially.")
	}
}
