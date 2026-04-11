package bitBlock

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
		blk.Set(data)

		expectedBuffer := []byte("Hello")
		if !bytes.Equal(blk.buffer, expectedBuffer) {
			t.Fatalf("Expected buffer %v, got %v", expectedBuffer, blk.buffer)
		}
	})

	// Sad path validation (nil input)
	t.Run("Sad path (nil input)", func(t *testing.T) {
		// Create a sample block with nil buffer
		blk := Block{}

		blk.Set(nil)

		if len(blk.buffer) != 0 {
			t.Fatalf("Expected 0-length buffer on nil input")
		}
	})

	// Sad path validation (nil buffer)
	t.Run("Sad path (nil buffer)", func(t *testing.T) {
		// Create a sample block with nil buffer and non-nil input
		blk := Block{}
		data := []byte("Hello")

		blk.Set(data)

		expectedBuffer := []byte("Hello")
		if !bytes.Equal(blk.buffer, expectedBuffer) {
			t.Fatalf("Expected buffer %v, got %v", expectedBuffer, blk.buffer)
		}
	})
}
