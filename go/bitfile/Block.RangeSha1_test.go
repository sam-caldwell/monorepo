package bitfile

/*
 * Block.RangeSha1() test
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * Test for the RangeSha1() method.
 */

import (
	"bytes"
	"testing"
)

func TestRangeSha1(t *testing.T) {
	// Create a sample block with some bytes
	blk := Block{
		buffer: []byte("Hello, World!"),
	}

	// Happy path validation
	t.Run("Happy path", func(t *testing.T) {
		start := 0
		stop := len(blk.buffer)

		hash, err := blk.RangeSha1(start, stop)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expectedHash := []byte{0x2ef7bde608ce5404e97d5f042f95f89f1c232871}
		if !bytes.Equal(hash, expectedHash) {
			t.Errorf("Expected hash %x, got %x", expectedHash, hash)
		}
	})

	// Sad path validations
	t.Run("Bounds check error (stop)", func(t *testing.T) {
		start := 0
		stop := len(blk.buffer) + 1

		_, err := blk.RangeSha1(start, stop)

		if err == nil || err.Error() != "bounds check error" {
			t.Errorf("Expected 'bounds check error', got %v", err)
		}
	})

	t.Run("Bounds check error (start)", func(t *testing.T) {
		start := -1
		stop := len(blk.buffer)

		_, err := blk.RangeSha1(start, stop)

		if err == nil || err.Error() != "bounds check error (start)" {
			t.Errorf("Expected 'bounds check error (start)', got %v", err)
		}
	})

	t.Run("Bounds check error (stop)", func(t *testing.T) {
		start := 0
		stop := -1

		_, err := blk.RangeSha1(start, stop)

		if err == nil || err.Error() != "bounds check error (stop)" {
			t.Errorf("Expected 'bounds check error (stop)', got %v", err)
		}
	})

	t.Run("Stop exceeds start", func(t *testing.T) {
		start := 5
		stop := 4

		_, err := blk.RangeSha1(start, stop)

		if err == nil || err.Error() != "stop exceeds start" {
			t.Errorf("Expected 'stop exceeds start', got %v", err)
		}
	})
}
