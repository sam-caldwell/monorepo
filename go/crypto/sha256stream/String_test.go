package crypto

import (
	"testing"
)

const (
	testString        = "test"
	expectedHashValue = "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
)

func TestSha256Stream_HashValue(t *testing.T) {
	// Create a new Sha256Stream instance
	hashStream := NewSha256Stream()

	// Convert the test string to a byte slice and add each byte as a bit to the stream
	for _, c := range testString {
		for i := 7; i >= 0; i-- {
			bit := ((c >> i) & 1) == 1
			hashStream.AddBit(bit)
		}
	}

	// Get the actual SHA256 hash as a hexadecimal string using the String method
	actualHashValue := hashStream.String()

	// Compare the actual and expected hash values
	if actualHashValue != expectedHashValue {
		t.Errorf("HashValue test failed. Expected: %s, Got: %s", expectedHashValue, actualHashValue)
	}
}
