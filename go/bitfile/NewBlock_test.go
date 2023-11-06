package bitfile

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
		size := 100
		blk := NewBlock(size)

		if len(blk.buffer) != size {
			t.Errorf("Expected buffer size %d, got %d", size, len(blk.buffer))
		}
	})

	// Sad path validation (negative size)
	t.Run("Sad path (negative size)", func(t *testing.T) {
		size := -1
		blk := NewBlock(size)

		if len(blk.buffer) != 0 {
			t.Errorf("Expected buffer size 0, got %d", len(blk.buffer))
		}
	})

	// Sad path validation (zero size)
	t.Run("Sad path (zero size)", func(t *testing.T) {
		size := 0
		blk := NewBlock(size)

		if len(blk.buffer) != size {
			t.Errorf("Expected buffer size %d, got %d", size, len(blk.buffer))
		}
	})
}
