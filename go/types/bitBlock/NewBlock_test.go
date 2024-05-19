package bitBlock

/*
 * NewBlock() Test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test the NewBlock() function.
 */

import (
	"testing"
)

func TestNewBlock(t *testing.T) {
	// Happy path validation
	t.Run("Happy path", func(t *testing.T) {
		size := uint(100)
		block := NewBlock(size)

		if len(block.buffer) != int(size) {
			t.Fatalf("Expected buffer size %d, got %d", size, len(block.buffer))
		}
	})

	// Sad path validation (zero size)
	t.Run("Sad path (zero size)", func(t *testing.T) {
		size := uint(0)
		blk := NewBlock(size)

		if len(blk.buffer) != int(size) {
			t.Fatalf("Expected buffer size %d, got %d", size, len(blk.buffer))
		}
	})
}
