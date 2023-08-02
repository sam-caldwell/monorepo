package crypto

import (
	"testing"
)

func TestNewSha256Stream(t *testing.T) {
	// Create a new Sha256Stream instance using NewSha256Stream function
	hash := NewSha256Stream()

	// Expected initial hash.h values (from SHA-256 specification)
	expectedH := [8]uint32{
		0x6a09e667,
		0xbb67ae85,
		0x3c6ef372,
		0xa54ff53a,
		0x510e527f,
		0x9b05688c,
		0x1f83d9ab,
		0x5be0cd19,
	}
	for i := 0; i < len(hash.h); i++ {
		if hash.h[i] != expectedH[i] {
			t.Fatalf("Expected h: %d Got: %d", expectedH[i], hash.h[i])
		}
	}
	if hash.bitNdx != 0 {
		t.Fatal("bitNdx expected 0")
	}
	if hash.byteNdx != 0 {
		t.Fatal("byteNdx expected 0")
	}
	if hash.size != 0 {
		t.Fatal("size expected 0")
	}
	for i := 0; i < len(hash.buffer); i++ {
		if hash.buffer[i] != 0x00 {
			t.Fatalf("expect empty buffer at i: %d", i)
		}
	}
}
