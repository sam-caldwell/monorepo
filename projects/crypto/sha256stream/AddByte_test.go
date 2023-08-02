package crypto

import (
	"testing"
)

func TestSha256Stream_AddByte_OneByte(t *testing.T) {
	// Create a new Sha256Stream instance
	hash := NewSha256Stream()
	if hash.bitNdx != 0 {
		t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
	}
	if hash.byteNdx != 0 {
		t.Fatalf("expected byteNdx 0. Got: %d", hash.bitNdx)
	}
	hash.AddByte(0xFF)
	if hash.bitNdx != 0 {
		t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
	}
	if hash.byteNdx != 1 {
		t.Fatalf("expected byteNdx 1. Got: %d", hash.bitNdx)
	}
}

func TestSha256Stream_AddByte_TwoBytes(t *testing.T) {
	// Create a new Sha256Stream instance
	hash := NewSha256Stream()
	if hash.bitNdx != 0 {
		t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
	}
	if hash.byteNdx != 0 {
		t.Fatalf("expected byteNdx 0. Got: %d", hash.bitNdx)
	}
	hash.AddByte(0xFF)
	if hash.bitNdx != 0 {
		t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
	}
	if hash.byteNdx != 1 {
		t.Fatalf("expected byteNdx 1. Got: %d", hash.bitNdx)
	}
	hash.AddByte(0xFF)
	if hash.bitNdx != 0 {
		t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
	}
	if hash.byteNdx != 2 {
		t.Fatalf("expected byteNdx 2. Got: %d", hash.bitNdx)
	}
}

func TestSha256Stream_AddByte_FullBuffer(t *testing.T) {
	// Create a new Sha256Stream instance
	hash := NewSha256Stream()
	for i := int8(1); i < 64; i++ {
		hash.AddByte(0xFF)
		if hash.bitNdx != 0 {
			t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
		}
		if hash.byteNdx != i {
			t.Fatalf("expected byteNdx %d. Got: %d", i, hash.bitNdx)
		}
	}
	hash.AddByte(0xFF)
	if hash.bitNdx != 0 {
		t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
	}
	if hash.byteNdx != 0 {
		t.Fatalf("expect byteNdx 0. Got: %d", hash.byteNdx)
	}
}
