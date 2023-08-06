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
	if hash.buffer != 0x03 {
		t.Fatal("expect 0b00000011 or 0x03")
	}
	hash.AddBit(false)
	if hash.buffer != 0x06 {
		t.Fatalf("expect 0x06")
	}
}

func TestSha256Stream_AddBit_Byte(t *testing.T) {
	hash := NewSha256Stream()
	hash.AddBit(true)  // 1
	hash.AddBit(false) // 2
	hash.AddBit(true)  // 3
	hash.AddBit(false) // 4
	hash.AddBit(true)  // 5
	hash.AddBit(false) // 6
	hash.AddBit(true)  // 7
	if hash.bitNdx != 7 {
		t.Fatalf("expected bitNdx 7. Got: %d", hash.bitNdx)
	}
	if hash.buffer != 0x55 {
		t.Fatalf("expected 0x55 in buffer. Got: %0x", hash.buffer)
	}
	hash.AddBit(false) // 8
	if hash.bitNdx != 0 {
		t.Fatalf("expected bitNdx 0. Got: %d", hash.bitNdx)
	}
	if hash.buffer != 0x00 {
		t.Fatalf("expected 0xAA in buffer. Got: %0x", hash.buffer)
	}
	hash.AddBit(true) // 1
	if hash.bitNdx != 1 {
		t.Fatalf("expected bitNdx 1. Got: %d", hash.bitNdx)
	}
	if hash.buffer != 0x01 {
		t.Fatal("expected 0x01 in buffer")
	}
}
