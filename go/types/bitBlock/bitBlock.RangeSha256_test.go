package bitBlock

/*
 * Block.RangeSha256() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test for the RangeSha256() method.
 */

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

func TestRangeSha256(t *testing.T) {
	block := Block{
		buffer: []byte("Hello, World!"),
	}

	// Happy path validation
	t.Run("Happy path", func(t *testing.T) {
		start := 0
		stop := len(block.buffer)

		hash, err := block.RangeSha256(start, stop)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expectedHash := sha256.Sum256(block.buffer[start:stop])
		if !bytes.Equal(hash, expectedHash[:]) {
			t.Errorf("Expected hash %x, got %x", expectedHash[:], hash)
		}
	})

	// Sad path validations
	t.Run("Bounds check error (stop)", func(t *testing.T) {
		start := 0
		stop := len(block.buffer) + 1

		_, err := block.RangeSha256(start, stop)

		if err == nil || err.Error() != "bounds check error" {
			t.Errorf("Expected 'bounds check error', got %v", err)
		}
	})

	t.Run("Bounds check error (start)", func(t *testing.T) {
		start := -1
		stop := len(block.buffer)

		_, err := block.RangeSha256(start, stop)

		if err == nil || err.Error() != "bounds check error (start)" {
			t.Errorf("Expected 'bounds check error (start)', got %v", err)
		}
	})

	t.Run("Bounds check error (stop)", func(t *testing.T) {
		start := 0
		stop := -1

		_, err := block.RangeSha256(start, stop)

		if err == nil || err.Error() != "bounds check error (stop)" {
			t.Errorf("Expected 'bounds check error (stop)', got %v", err)
		}
	})

	t.Run("Stop exceeds start", func(t *testing.T) {
		start := 5
		stop := 4

		_, err := block.RangeSha256(start, stop)

		if err == nil || err.Error() != "stop exceeds start" {
			t.Errorf("Expected 'stop exceeds start', got %v", err)
		}
	})
}
