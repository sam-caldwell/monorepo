package crypto

import (
	"testing"
)

func TestSha256StreamClearBuffer(t *testing.T) {
	// Create a new Sha256Stream instance
	stream := Sha256Stream{}

	// Fill the buffer with some arbitrary non-zero values
	for i := 0; i < 64; i++ {
		stream.buffer[i] = byte(i) + 1
	}

	// Set bitNdx and byteNdx to some arbitrary non-zero values
	stream.bitNdx = 3
	stream.byteNdx = 2

	// Call the ClearBuffer method
	stream.ClearBuffer(false)

	// Check if the buffer has been cleared (all elements are 0) after calling ClearBuffer
	for i := 0; i < 64; i++ {
		if stream.buffer[i] != 0 {
			t.Errorf("ClearBuffer() method didn't clear the buffer as expected.\nExpected buffer[%d]: 0\nActual buffer[%d]: %d", i, i, stream.buffer[i])
		}
	}

	// Check if bitNdx and byteNdx have been reset to 0 after calling ClearBuffer
	if stream.bitNdx != 0 {
		t.Errorf("ClearBuffer() method didn't reset bitNdx as expected.\nExpected bitNdx: 0\nActual bitNdx: %d", stream.bitNdx)
	}

	if stream.byteNdx != 0 {
		t.Errorf("ClearBuffer() method didn't reset byteNdx as expected.\nExpected byteNdx: 0\nActual byteNdx: %d", stream.byteNdx)
	}
}
