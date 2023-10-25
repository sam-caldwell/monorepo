package bitfile

/*
 * CRSCE BitFile
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * bit-for-bit reader/writer
 */

import (
	"testing"
)

// TestBitFile_Struct - test the structure itself.
func TestBitFile_Struct(t *testing.T) {
	var b BitFile

	if b.file != nil {
		t.Fatal("Bitfile.file should be nil initially.")
	}
	if b.filePos != 0 {
		t.Fatal("Bitfile.filePos should be 0 initially.")
	}
	if b.bitPos != 0 {
		t.Fatal("Bitfile.bitPos should be 0 initially.")
	}
	if b.bufferPos != 0 {
		t.Fatal("Bitfile.bufferPos should be 0 initially.")
	}
	if b.bufferSize != 0 {
		t.Fatalf("Bitfile.bufferSize should be %d initially", 0)
	}
	if bufferSize < 1024 {
		t.Fatal("Bitfile should have a constant bufferSize of at least 1024")
	}
}
