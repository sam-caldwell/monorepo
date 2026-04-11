package crypto

import (
	"testing"
)

func TestSha256StreamSize_initialState(t *testing.T) {
	// Create a new Sha256Stream instance
	hash := NewSha256Stream()
	if hash.Size() != 32 {
		t.Fatal("initial size 0")
	}
}

func TestSha256StreamSize_oneBit(t *testing.T) {
	// Create a new Sha256Stream instance
	hash := NewSha256Stream()
	hash.AddBit(true)
	if hash.Size() != 32 {
		t.Fatal("initial size 0")
	}
}
