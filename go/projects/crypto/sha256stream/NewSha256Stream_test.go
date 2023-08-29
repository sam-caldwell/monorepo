package crypto

import (
	"testing"
)

func TestNewSha256Stream(t *testing.T) {
	// Create a new Sha256Stream instance using NewSha256Stream function
	hash := NewSha256Stream()
	if hash.bitNdx != 0 {
		t.Fatal("bitNdx expected 0")
	}
	if hash.buffer != 0x00 {
		t.Fatal("buffer should be a null byte")
	}
}
