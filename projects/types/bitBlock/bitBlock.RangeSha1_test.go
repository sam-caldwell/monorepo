package bitBlock

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
	block := Block{
		buffer: []byte("Hello, World!"),
	}

	expectedHash := []byte{
		0x0a, 0x0a, 0x9f, 0x2a, 0x67, 0x72, 0x94, 0x25, 0x57, 0xab,
		0x53, 0x55, 0xd7, 0x6a, 0xf4, 0x42, 0xf8, 0xf6, 0x5e, 0x01}

	// Happy path validation
	t.Run("Happy path", func(t *testing.T) {
		var err error
		var actual []byte
		start := 0
		stop := len(block.buffer)

		if actual, err = block.RangeSha1(start, stop); err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if !bytes.Equal(actual, expectedHash) {
			t.Fatalf("Expected hash %x, got %x", expectedHash, actual)
		}
	})

	// Sad path validations
	t.Run("Bounds check error (stop)", func(t *testing.T) {
		var err error
		start := 0
		stop := len(block.buffer) + 1

		_, err = block.RangeSha1(start, stop)
		if err == nil || err.Error() != "bounds check error" {
			t.Fatalf("Expected 'bounds check error', got %v", err)
		}
	})

	t.Run("Bounds check error (start)", func(t *testing.T) {
		var err error
		start := -1
		stop := len(block.buffer)

		_, err = block.RangeSha1(start, stop)
		if err == nil || err.Error() != "bounds check error (start)" {
			t.Fatalf("Expected 'bounds check error (start)', got %v", err)
		}
	})

	t.Run("Bounds check error (stop)", func(t *testing.T) {
		var err error
		start := 0
		stop := -1

		_, err = block.RangeSha1(start, stop)
		if err == nil || err.Error() != "bounds check error (stop)" {
			t.Fatalf("Expected 'bounds check error (stop)', got %v", err)
		}
	})

	t.Run("Stop exceeds start", func(t *testing.T) {
		var err error
		start := 5
		stop := 4

		_, err = block.RangeSha1(start, stop)
		if err == nil || err.Error() != "stop exceeds start" {
			t.Fatalf("Expected 'stop exceeds start', got %v", err)
		}
	})
}
