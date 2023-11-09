package bitfile

/*
 * bitfile tests
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * Test the bitfile struct
 */

import (
	"testing"
)

// TestBitFile_Struct - test the structure itself.
func TestBitFile_Struct(t *testing.T) {
	var b Reader

	//Validate the constant used for default buffer size
	if defaultBufferSize < 4*1048576 {
		t.Fatal("Bitfile should have a constant defaultBufferSize of at least 1024")
	}

	if b.file != nil {
		t.Fatal("Bitfile.file should be nil initially.")
	}
}
