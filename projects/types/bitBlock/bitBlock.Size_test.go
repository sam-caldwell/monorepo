package bitBlock

import "testing"

/*
 * bitBlock.Size() method unit test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Unit test for the size() method.
 */

func TestBlock_Size(t *testing.T) {
	var block Block

	if block.Size() != 0 {
		t.Fatal("block size 0 initially expected")
	}

	block.buffer = make([]byte, 10)

	if block.Size() != 10 {
		t.Fatal("block size mismatch")
	}
}
