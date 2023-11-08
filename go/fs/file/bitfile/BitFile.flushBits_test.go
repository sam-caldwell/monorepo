package file

import "testing"

/*
 * CRSCE BitFile.flushBits() Test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * When bits are written to the BitFile they are first buffered to a single byte.  Every 8th bit
 * can then be written to a byte-level buffer which is ultimately flushed to disk.  This file tests
 * that functionality.
 */

func TestBitFile_flushBits(t *testing.T) {
	t.Skip("not implemented")
}
