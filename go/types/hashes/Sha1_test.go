package hashes

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestStateMetric_StateType_Sha1(t *testing.T) {
	const expectedSize = uint(20)
	/*
	 * Test Sha256Length
	 */
	t.Run("Test Sha1Length", func(t *testing.T) {
		if expectedSize != Sha1Length {
			t.Fatal("length is wrong")
		}
	})
	/*
	 * Tests for FromSlice()
	 */
	t.Run("Test FromSlice() Happy", func(t *testing.T) {
		var hash Sha1
		actual := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
		}
		expected := [20]byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
		}
		hash.FromSlice(actual)
		if !bytes.Equal(actual, expected[:]) {
			t.Fatalf("Mismatch\n"+
				"Actual:  %v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	t.Run("Test FromSlice() Truncating overflow", func(t *testing.T) {
		var hash Sha1
		actual := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, //These are overflow
		}
		expected := [20]byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
		}
		hash.FromSlice(actual)
		if !bytes.Equal(actual[:20], expected[:]) {
			t.Fatalf("Mismatch\n"+
				"Actual:  %v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	t.Run("Test FromSlice() Empty Slice", func(t *testing.T) {
		var hash Sha1
		actual := []byte{
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
		}
		var expected [20]byte //leave it empty
		hash.FromSlice([]byte{})
		if !bytes.Equal(actual[:], expected[:]) {
			t.Fatalf("Mismatch\n"+
				"Actual:  %v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	t.Run("Test FromSlice() nil Slice", func(t *testing.T) {
		var hash Sha1
		actual := []byte{
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
		}
		var expected [20]byte //leave it empty
		hash.FromSlice(nil)
		if !bytes.Equal(actual[:], expected[:]) {
			t.Fatalf("Mismatch\n"+
				"Actual:  %v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	/*
	 * Tests for ToSlice()
	 */
	t.Run("Test ToSlice() Happy", func(t *testing.T) {
		var hash Sha1
		expected := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
		}
		hash.FromSlice(expected)
		actual := hash.ToSlice()
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Mismatch\n"+
				"Actual:  %v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	t.Run("Test ToSlice() Empty Slice", func(t *testing.T) {
		var hash Sha1
		expected := []byte{
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
		}
		hash.FromSlice([]byte{})
		actual := hash.ToSlice()
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Mismatch\n"+
				"Actual:  %v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	/*
	 * Test for SizeOf()
	 */
	t.Run("Test ToSlice() Happy 1", func(t *testing.T) {
		var hash Sha1
		if hash.SizeOf() != expectedSize {
			t.Fatalf("Size Mismatch")
		}
	})
	t.Run("Test ToSlice() Happy 2", func(t *testing.T) {
		var hash Sha1
		expected := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
		}
		hash.FromSlice(expected)
		if hash.SizeOf() != expectedSize {
			t.Fatalf("Size Mismatch")
		}
	})
	t.Run("Test HashString() Happy", func(t *testing.T) {
		const testInput = "This is my test.  There is no test like my test."
		const expected = "027c4109f56c0dffdcfd2e0dc50fdc97488d5300"
		var hash Sha1
		hash.HashString(testInput)
		result := hash.ToSlice()
		if actual := hex.EncodeToString(result); actual != expected {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	t.Run("Test String() Happy", func(t *testing.T) {
		const testInput = "This is my test.  There is no test like my test."
		const expected = "027c4109f56c0dffdcfd2e0dc50fdc97488d5300"
		var hash Sha1
		hash.HashString(testInput)
		if actual := hash.String(); actual != expected {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
}
