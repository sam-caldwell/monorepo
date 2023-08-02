package crypto

import (
	"testing"
)

func TestSha256StreamStructure(t *testing.T) {
	stream := Sha256Stream{}

	// Test the initial values of the struct fields
	if stream.byteNdx != 0 {
		t.Errorf("Expected byteNdx to be 0, but got %d", stream.byteNdx)
	}

	if stream.bitNdx != 0 {
		t.Errorf("Expected bitNdx to be 0, but got %d", stream.bitNdx)
	}

	if stream.size != 0 {
		t.Errorf("Expected size to be 0, but got %d", stream.size)
	}

	// Test the length of the h array
	if len(stream.h) != 8 {
		t.Errorf("Expected length of h array to be 8, but got %d", len(stream.h))
	}

	// Test the length of the buffer array
	if len(stream.buffer) != 64 {
		t.Errorf("Expected length of buffer array to be 64, but got %d", len(stream.buffer))
	}
}
