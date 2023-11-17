package generic

import (
	"encoding/hex"
	"reflect"
	"strings"
	"testing"
)

//goland:noinspection GoBoolExpressions
func TestBlock2KB(t *testing.T) {
	var b Block2KB
	if len(b) != 2048 {
		t.Fatalf("Size error (%d)", len(b))
	}
}

func TestBlock2KB_FromSlice(t *testing.T) {
	// Happy path
	t.Run("HappyPath", func(t *testing.T) {
		var block Block2KB

		testInput := []byte("This is a test block with size.")

		block.FromSlice(testInput)
		expected := make([]byte, len(block))
		copy(expected, testInput)

		if !reflect.DeepEqual(block[:], expected) {
			t.Errorf("data mismatch.\n"+
				"Actual:   %v (%d)\n"+
				"Expected: %v (%d)",
				block[:], len(block[:]), expected, len(expected))
		}
	})

	// Sad path
	t.Run("SadPath_InsufficientData", func(t *testing.T) {
		var block Block2KB

		// Test with a slice that has insufficient data
		insufficientData := []byte("Short")

		block.FromSlice(insufficientData)

		// Ensure that only the available data is copied, and the rest is zero-filled
		expectedData := append(insufficientData, make([]byte, len(block)-len(insufficientData))...)
		if !reflect.DeepEqual(block[:], expectedData) {
			t.Errorf("data mismatch. Expected: %v, Got: %v", expectedData, block[:])
		}
	})

	// Sad path
	t.Run("SadPath_ExcessiveData", func(t *testing.T) {

		var block Block2KB

		// Test with a slice that has excessive data
		excessiveData := make([]byte, 2*len(block))

		block.FromSlice(excessiveData)

		// Ensure that only the first len(block) bytes are copied
		expectedData := excessiveData[:len(block)]
		if !reflect.DeepEqual(block[:], expectedData) {
			t.Errorf("data mismatch. Expected: %v, Got: %v", expectedData, block[:])
		}
	})
}
func TestBlock2KB_SizeOf(t *testing.T) {
	const blockSize = 2048
	// Happy path
	t.Run("HappyPath", func(t *testing.T) {
		var block Block2KB

		size := block.SizeOf()

		expectedSize := uint(blockSize)
		if size != expectedSize {
			t.Errorf("size mismatch. Expected: %d, Got: %d", expectedSize, size)
		}
	})

	// Sad path
	t.Run("SadPath_ModifyBlock", func(t *testing.T) {
		// Modify the block, and then check the size
		var block Block2KB
		block[0] = 42 // Modify the block

		size := block.SizeOf()

		// The size should still be 'blockSize', as modifying the contents doesn't change the size
		expectedSize := uint(blockSize)
		if size != expectedSize {
			t.Errorf("size mismatch. Expected: %d, Got: %d", expectedSize, size)
		}
	})
}

func TestBlock2KB_String(t *testing.T) {
	const blockSize = 2048
	// Happy path
	t.Run("HappyPath", func(t *testing.T) {
		var block Block2KB

		// Assuming block is filled with zeros in the happy path
		testInput := "" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
			"123456789123456789123-5678" + strings.Repeat("00", blockSize-1024)

		block.FromSlice([]byte(testInput))
		result := block.String()

		expectedString := hex.EncodeToString([]byte(testInput)[:blockSize])

		if result != expectedString {
			t.Errorf("String() mismatch.\n"+
				"Expected: %s\n"+
				"Got: %s",
				expectedString, result)
		}
	})

	// Sad path
	t.Run("SadPath_NonZeroBlock", func(t *testing.T) {
		// Modify the block, and then check the string representation
		var block Block2KB
		block[0] = 42 // Modify the block

		expectedString := "2a" + strings.Repeat("00", blockSize-1)

		result := block.String()

		if result != expectedString {
			t.Errorf("String() mismatch.\n"+
				"Expected: %s\n"+
				"Got: %s",
				expectedString, result)
		}
	})
}

func TestBlock2KB_ToSlice(t *testing.T) {
	const blockSize = 2048
	// Happy path
	t.Run("HappyPath", func(t *testing.T) {
		var block Block2KB
		// Assuming block is filled with zeros in the happy path
		expectedSlice := make([]byte, blockSize)

		result := block.ToSlice()

		if !reflect.DeepEqual(result, expectedSlice) {
			t.Errorf("Block1KB ToSlice() mismatch. Expected: %v, Got: %v", expectedSlice, result)
		}
	})

	// Sad path
	t.Run("SadPath_NonZeroBlock", func(t *testing.T) {
		// Modify the block, and then check the byte slice representation
		var block Block2KB
		block[0] = 42 // Modify the block

		// The expected slice should have the modified byte and zeros for the rest
		expectedSlice := make([]byte, blockSize)
		expectedSlice[0] = 42

		result := block.ToSlice()

		if !reflect.DeepEqual(result, expectedSlice) {
			t.Errorf("Block1KB ToSlice() mismatch. Expected: %v, Got: %v", expectedSlice, result)
		}
	})
}
