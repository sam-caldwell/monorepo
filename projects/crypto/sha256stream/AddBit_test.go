package crypto

import (
	"testing"
)

func TestSha256Stream_singleBit(t *testing.T) {
	hash := NewSha256Stream()
	hash.AddBit(true)
	if hash.bitNdx != 1 {
		t.Fatal("expected bitNdx 1")
	}
	hash.AddBit(true)
	if hash.bitNdx != 2 {
		t.Fatal("expected bitNdx 2")
	}
}

func TestSha256Stream_AddBit_ByteFlushed(t *testing.T) {
	hash := NewSha256Stream()
	hash.AddBit(true) // 1
	hash.AddBit(true) // 2
	hash.AddBit(true) // 3
	hash.AddBit(true) // 4
	hash.AddBit(true) // 5
	hash.AddBit(true) // 6
	hash.AddBit(true) // 7
	hash.AddBit(true) // 8
	if hash.bitNdx != 0 {
		t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
	}
	if hash.byteNdx != 1 {
		t.Fatalf("expect byteNdx 1. Got: %d", hash.byteNdx)
	}
	hash.AddBit(true) // 1
	if hash.bitNdx != 1 {
		t.Fatalf("expected bitNdx 1. Got: %d", hash.bitNdx)
	}
	if hash.byteNdx != 1 {
		t.Fatalf("expect byteNdx 1. Got: %d", hash.byteNdx)
	}
}

func TestSha256Stream_AddBit_FullBuffer(t *testing.T) {
	hash := NewSha256Stream()
	for i := 0; i < 511; i++ {
		hash.AddBit(true)
	}
	if hash.bitNdx != 7 {
		t.Fatalf("expected bitNdx 7. Got: %d", hash.bitNdx)
	}
	if hash.byteNdx != int8(len(hash.buffer)-1) {
		t.Fatalf("expect byteNdx %d. Got: %d", len(hash.buffer)-1, hash.byteNdx)
	}
	hash.AddBit(true)
	if hash.bitNdx != 0 {
		t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
	}
	if hash.byteNdx != 0 {
		t.Fatalf("expect byteNdx 0. Got: %d", hash.byteNdx)
	}
}
