package bitBlock

import (
	"bytes"
	"testing"
)

/*
 * bitBlock.ReadBytes() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test the ReadBytes() method.
 */

func TestBlock_ReadBytes(t *testing.T) {
	expected := []byte("This is a test message")
	var block Block
	block.buffer = make([]byte, 1024)
	block.buffer = expected
	if actual := block.ReadBytes(0); !bytes.Equal(actual, expected) {
		t.Fatal("ReadBytes mismatch")
	}
	if actual := block.ReadBytes(10); !bytes.Equal(actual, expected[0:10]) {
		t.Fatal("ReadBytes mismatch (10)")
	}
}
