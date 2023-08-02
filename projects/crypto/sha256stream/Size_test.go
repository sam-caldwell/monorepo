package crypto

import (
	"testing"
)

func TestSha256StreamSize_initialState(t *testing.T) {
	// Create a new Sha256Stream instance
	hash := NewSha256Stream()
	if hash.Size() != 0 {
		t.Fatal("initial size 0")
	}
}

func TestSha256StreamSize_oneBit(t *testing.T) {
	// Create a new Sha256Stream instance
	hash := NewSha256Stream()
	hash.AddBit(true)
	if hash.Size() != 0 {
		t.Fatal("initial size 0")
	}
}

func TestSha256StreamSize_oneByte(t *testing.T) {
	// Create a new Sha256Stream instance
	hash := NewSha256Stream()
	hash.AddByte(0xFF)
	if hash.Size() != 0 {
		t.Fatal("initial size 0")
	}
}

func TestSha256StreamSize_oneFullBuffer(t *testing.T) {
	// Create a new Sha256Stream instance
	hash := NewSha256Stream()
	for i := int8(0); i < 64; i++ {
		hash.AddByte(0xFF)
		if hash.bitNdx != 0 {
			t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
		}
	}
	if hash.byteNdx != 0 {
		t.Fatalf("expected byteNdx 1. Got: %d", hash.byteNdx)
	}
	if hash.Size() != len(hash.buffer) {
		t.Fatalf("initial size Got: %d\n"+
			"Expected:%d",
			hash.Size(),
			uint32(len(hash.buffer)))
	}
}
