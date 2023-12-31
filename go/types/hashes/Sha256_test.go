package hashes

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestStateMetric_StateType_Sha256(t *testing.T) {
	const expectedSize = uint(32)
	/*
	 * Test Sha256Length
	 */
	t.Run("Test Sha256Length", func(t *testing.T) {
		if expectedSize != Sha256Length {
			t.Fatal("length is wrong")
		}
	})
	/*
	 * Tests for FromSlice()
	 */
	t.Run("Test FromSlice() Happy", func(t *testing.T) {
		var hash Sha256
		actual := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F,
		}
		expected := [expectedSize]byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F,
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
		var hash Sha256
		actual := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, //These are overflow
		}
		expected := [expectedSize]byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F,
		}
		hash.FromSlice(actual)
		if !bytes.Equal(actual[:expectedSize], expected[:]) {
			t.Fatalf("Mismatch\n"+
				"Actual:  %v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	t.Run("Test FromSlice() Empty Slice", func(t *testing.T) {
		var hash Sha256
		actual := []byte{
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00,
		}
		var expected [expectedSize]byte //leave it empty
		hash.FromSlice([]byte{})
		if !bytes.Equal(actual[:], expected[:]) {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	t.Run("Test FromSlice() nil Slice", func(t *testing.T) {
		var hash Sha256
		actual := []byte{
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00,
		}
		var expected [expectedSize]byte //leave it empty
		hash.FromSlice(nil)
		if !bytes.Equal(actual[:], expected[:]) {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	/*
	 * Tests for ToSlice()
	 */
	t.Run("Test ToSlice() Happy", func(t *testing.T) {
		var hash Sha256
		expected := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F,
		}
		hash.FromSlice(expected)
		actual := hash.ToSlice()
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Mismatch\n"+
				"Actual:%v (%d)\n"+
				"Expected:%v (%d)\n",
				actual, len(actual), expected, len(expected))
		}
	})
	t.Run("Test ToSlice() Empty Slice", func(t *testing.T) {
		var hash Sha256
		expected := []byte{
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00,
		}
		hash.FromSlice([]byte{})
		actual := hash.ToSlice()
		if !bytes.Equal(actual, expected) {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	/*
	 * Test for SizeOf()
	 */
	t.Run("Test ToSlice() Happy 1", func(t *testing.T) {
		var hash Sha256
		if hash.SizeOf() != expectedSize {
			t.Fatalf("Size Mismatch")
		}
	})
	t.Run("Test ToSlice() Happy 2", func(t *testing.T) {
		var hash Sha256
		expected := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F,
		}
		hash.FromSlice(expected)
		if hash.SizeOf() != expectedSize {
			t.Fatalf("Size Mismatch")
		}
	})
	t.Run("Test ToSlice() Empty Slice", func(t *testing.T) {
		var hash Sha256
		hash.FromSlice([]byte{})
		if hash.SizeOf() != expectedSize {
			t.Fatalf("Size Mismatch")
		}
	})

	t.Run("Test HashBytes() Happy", func(t *testing.T) {
		const testInput = "This is my test.  There is no test like my test."
		const expected = "6e5d6a1041fe8a13fe5f8062e1c8d2741d71aae225b46c2efc539230f2113d25"
		var hash Sha256
		hash.HashBytes([]byte(testInput))
		actual := hash.ToSlice()
		if hex.EncodeToString(actual) != expected {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
	t.Run("Test HashString() Happy", func(t *testing.T) {
		const testInput = "This is my test.  There is no test like my test."
		const expected = "6e5d6a1041fe8a13fe5f8062e1c8d2741d71aae225b46c2efc539230f2113d25"
		var hash Sha256
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
		const expected = "6e5d6a1041fe8a13fe5f8062e1c8d2741d71aae225b46c2efc539230f2113d25"
		var hash Sha256
		hash.HashString(testInput)
		if actual := hash.String(); actual != expected {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
}
