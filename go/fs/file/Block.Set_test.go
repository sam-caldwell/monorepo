package crsce

/*
 * Block.Set() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test for the Set() method.
 */

import (
	"bytes"
	"testing"
)

func TestSet(t *testing.T) {
	// Happy path validation
	t.Run("Happy path", func(t *testing.T) {
		// Create a sample block
		blk := Block{}

		data := []byte("Hello")
		err := blk.Set(data)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expectedBuffer := []byte("Hello")
		if !bytes.Equal(blk.buffer, expectedBuffer) {
			t.Errorf("Expected buffer %v, got %v", expectedBuffer, blk.buffer)
		}
	})

	// Sad path validation (nil input)
	t.Run("Sad path (nil input)", func(t *testing.T) {
		// Create a sample block with nil buffer
		blk := Block{}

		err := blk.Set(nil)

		if err == nil || err.Error() != "invalid nil input" {
			t.Errorf("Expected 'invalid nil input' error, got %v", err)
		}
	})

	// Sad path validation (nil buffer)
	t.Run("Sad path (nil buffer)", func(t *testing.T) {
		// Create a sample block with nil buffer and non-nil input
		blk := Block{}
		data := []byte("Hello")

		err := blk.Set(data)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expectedBuffer := []byte("Hello")
		if !bytes.Equal(blk.buffer, expectedBuffer) {
			t.Errorf("Expected buffer %v, got %v", expectedBuffer, blk.buffer)
		}
	})
}
