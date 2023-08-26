package crypto

import (
	"crypto/sha256"
	"testing"
)

func TestSha256Stream_Output_Empty(t *testing.T) {
	// Create a new Sha256Stream instance
	hashStream := NewSha256Stream()

	// Get the expected SHA256 hash of the added bits
	expectedHash := sha256.Sum256([]byte{}) // Expected hash of 0b101 (binary) or 0x5 (hex)

	// Get the actual SHA256 hash using the Output method
	actualHash := hashStream.Output()

	// Compare the actual and expected hashes
	if !bytesEqual(actualHash, expectedHash[:]) {
		t.Errorf("Output test failed. Expected: %x, Got: %x", expectedHash[:], actualHash)
	}
}
func TestSha256Stream_Output_NonEmpty(t *testing.T) {
	// Create a new Sha256Stream instance
	hashStream := NewSha256Stream()

	// Add bits to the stream
	hashStream.AddBit(true)
	hashStream.AddBit(false)
	hashStream.AddBit(true)
	hashStream.AddBit(true)
	hashStream.AddBit(false)
	hashStream.AddBit(false)
	hashStream.AddBit(false)
	hashStream.AddBit(true)

	// Expected hash of the added bits: 0b10110001 (binary) or 0xB1 (hex)
	expectedHash := sha256.Sum256([]byte{0xB1})

	// Get the actual SHA256 hash using the Output method
	actualHash := hashStream.Output()

	// Compare the actual and expected hashes
	if !bytesEqual(actualHash, expectedHash[:]) {
		t.Errorf("Output test failed. Expected: %x, Got: %x", expectedHash[:], actualHash)
	}
}

// Helper function to compare two byte slices
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
