package bitfile

/*
 * Writer.Close() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Unit test for the Writer.Close() method
 */

import "testing"

func TestWriter_Struct(t *testing.T) {
	var w Writer

	if w.file != nil {
		t.Fatal("Writer expects to be nil initially")
	}

	if w.buffer != nil {
		t.Fatal("Writer expects nil buffer initially")
	}

	if w.bufferPos != 0 {
		t.Fatal("bufferPos should be 0")
	}

	if w.bitPos != 0 {
		t.Fatal("bitPos should be 0")
	}

}
