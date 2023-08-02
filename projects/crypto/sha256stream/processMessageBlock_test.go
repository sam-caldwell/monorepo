package crypto

import "testing"

func TestSha256StreamProcessMessageBlock(t *testing.T) {
	// Create a new Sha256Stream instance
	hash := NewSha256Stream()
	hash.
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).
		AddByte(0xFF).AddByte(0xFF).AddByte(0xFF).AddByte(0xFF)

	expectedState := []uint32{
		0xef0c748d,
		0xf4da50a8,
		0xd6c43c01,
		0x3edc3ce7,
		0x6c9d9fa9,
		0xa1458ade,
		0x56eb86c0,
		0xa64492d2,
	}

	if hash.Size() != len(hash.buffer) {
		t.Fatalf("expected buffer flush did not occur")
	}
	for i := 0; i < len(hash.h); i++ {
		if hash.h[i] != expectedState[i] {
			t.Errorf("unexpected hash state (%0d).\n"+
				"Expected: 0x%0x\n"+
				"Got:0x%0x", i, expectedState[i], hash.h[i])
		}
	}
}
