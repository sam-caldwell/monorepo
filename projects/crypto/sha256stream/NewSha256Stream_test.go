package crypto

import (
	"testing"
)

func TestNewSha256Stream(t *testing.T) {
	// Create a new Sha256Stream instance using NewSha256Stream function
	stream := NewSha256Stream()

	// Expected initial hash.h values (from SHA-256 specification)
	expectedH := [8]uint32{
		0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a,
		0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19,
	}

	// Check if the hash.h values have been initialized correctly
	for i, h := range expectedH {
		if stream.h[i] != h {
			t.Errorf("NewSha256Stream() didn't initialize hash.h correctly.\n"+
				"Expected hash.h[%d]: %x\n"+
				"Actual hash.h[%d]: %x", i, h, i, stream.h[i])
		}
	}

	// Check if the buffer, bitNdx, byteNdx, and size have been initialized correctly
	if len(stream.buffer) != 64 {
		t.Errorf("NewSha256Stream() didn't initialize buffer correctly.\n"+
			"Expected buffer length: %d\n"+
			"Actual buffer length: %d", 64, len(stream.buffer))
	}
	if stream.bitNdx != 0 {
		t.Errorf("NewSha256Stream() didn't initialize bitNdx correctly.\n"+
			"Expected bitNdx: %d\n"+
			"Actual bitNdx: %d", 0, stream.bitNdx)
	}
	if stream.byteNdx != 0 {
		t.Errorf("NewSha256Stream() didn't initialize byteNdx correctly.\n"+
			"Expected byteNdx: %d\n"+
			"Actual byteNdx: %d", 0, stream.byteNdx)
	}
	if stream.size != 0 {
		t.Errorf("NewSha256Stream() didn't initialize size correctly.\n"+
			"Expected size: %d\n"+
			"Actual size: %d", 0, stream.size)
	}
}
