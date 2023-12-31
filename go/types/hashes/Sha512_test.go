package hashes

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestStateMetric_StateType_Sha512(t *testing.T) {
	const expectedSize = uint(64)
	/*
	 * Test Sha256Length
	 */
	t.Run("Test Sha256Length", func(t *testing.T) {
		if expectedSize != Sha512Length {
			t.Fatal("length is wrong")
		}
	})
	/*
	 * Tests for FromSlice()
	 */
	t.Run("Test FromSlice() Happy", func(t *testing.T) {
		var hash Sha512
		actual := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F, 0x20, 0x27, 0x28,
			0x29, 0x2A, 0x2B, 0x2C, 0x2D,
			0x2E, 0x2F, 0x30, 0x31, 0x32,
			0x33, 0x34, 0x35, 0x36, 0x37,
			0x38, 0x39, 0x3A, 0x3B, 0x3C,
			0x3E, 0x3F, 0x40, 0x41, 0x42,
			0x43, 0x44, 0x45, 0x46,
		}
		expected := [expectedSize]byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F, 0x20, 0x27, 0x28,
			0x29, 0x2A, 0x2B, 0x2C, 0x2D,
			0x2E, 0x2F, 0x30, 0x31, 0x32,
			0x33, 0x34, 0x35, 0x36, 0x37,
			0x38, 0x39, 0x3A, 0x3B, 0x3C,
			0x3E, 0x3F, 0x40, 0x41, 0x42,
			0x43, 0x44, 0x45, 0x46,
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
		var hash Sha512
		actual := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F, 0x20, 0x27, 0x28,
			0x29, 0x2A, 0x2B, 0x2C, 0x2D,
			0x2E, 0x2F, 0x30, 0x31, 0x32,
			0x33, 0x34, 0x35, 0x36, 0x37,
			0x38, 0x39, 0x3A, 0x3B, 0x3C,
			0x3E, 0x3F, 0x40, 0x41, 0x42,
			0x43, 0x44, 0x45, 0x46, 0xFF,
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
			0x1E, 0x1F, 0x20, 0x27, 0x28,
			0x29, 0x2A, 0x2B, 0x2C, 0x2D,
			0x2E, 0x2F, 0x30, 0x31, 0x32,
			0x33, 0x34, 0x35, 0x36, 0x37,
			0x38, 0x39, 0x3A, 0x3B, 0x3C,
			0x3E, 0x3F, 0x40, 0x41, 0x42,
			0x43, 0x44, 0x45, 0x46,
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
		var hash Sha512
		actual := []byte{
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
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
		var hash Sha512
		actual := []byte{
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
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
		var hash Sha512
		expected := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F, 0x20, 0x27, 0x28,
			0x29, 0x2A, 0x2B, 0x2C, 0x2D,
			0x2E, 0x2F, 0x30, 0x31, 0x32,
			0x33, 0x34, 0x35, 0x36, 0x37,
			0x38, 0x39, 0x3A, 0x3B, 0x3C,
			0x3E, 0x3F, 0x40, 0x41, 0x42,
			0x43, 0x44, 0x45, 0x46,
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
		var hash Sha512
		expected := []byte{
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
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
		var hash Sha512
		if hash.SizeOf() != expectedSize {
			t.Fatalf("Size Mismatch")
		}
	})
	t.Run("Test ToSlice() Happy 2", func(t *testing.T) {
		var hash Sha512
		expected := []byte{
			0x00, 0x01, 0x02, 0x03, 0x04,
			0x05, 0x06, 0x07, 0x08, 0x09,
			0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
			0x0F, 0x10, 0x11, 0x12, 0x13,
			0x14, 0x15, 0x16, 0x17, 0x18,
			0x19, 0x1A, 0x1B, 0x1C, 0x1D,
			0x1E, 0x1F, 0x20, 0x27, 0x28,
			0x29, 0x2A, 0x2B, 0x2C, 0x2D,
			0x2E, 0x2F, 0x30, 0x31, 0x32,
			0x33, 0x34, 0x35, 0x36, 0x37,
			0x38, 0x39, 0x3A, 0x3B, 0x3C,
			0x3E, 0x3F, 0x40, 0x41, 0x42,
			0x43, 0x44, 0x45, 0x46,
		}
		hash.FromSlice(expected)
		if hash.SizeOf() != expectedSize {
			t.Fatalf("Size Mismatch")
		}
	})
	t.Run("Test ToSlice() Empty Slice", func(t *testing.T) {
		var hash Sha512
		hash.FromSlice([]byte{})
		if hash.SizeOf() != expectedSize {
			t.Fatalf("Size Mismatch")
		}
	})
	t.Run("Test HashString() Happy", func(t *testing.T) {
		const testInput = "This is my test.  There is no test like my test."
		const expected = "a7096bad62963e64706aebbbf0558d4435373a5a98ee58200f29674938ccff461403a1c76febb3ebb88eddbf642a04d65bf328f97d9fab807afc5d88088f1b75"
		var hash Sha512
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
		const expected = "a7096bad62963e64706aebbbf0558d4435373a5a98ee58200f29674938ccff461403a1c76febb3ebb88eddbf642a04d65bf328f97d9fab807afc5d88088f1b75"
		var hash Sha512
		hash.HashString(testInput)
		if actual := hash.String(); actual != expected {
			t.Fatalf("Mismatch\n"+
				"Actual:%v\n"+
				"Expected:%v\n",
				actual, expected)
		}
	})
}
