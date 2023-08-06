package crypto

import (
	"testing"
)

func TestSha256StreamStructure(t *testing.T) {
	stream := Sha256Stream{}
	if stream.bitNdx != 0 {
		t.Errorf("Expected bitNdx to be 0, but got %d", stream.bitNdx)
	}
	// Test the length of the buffer array
	if stream.buffer != 0x00 {
		t.Fail()
	}
}
