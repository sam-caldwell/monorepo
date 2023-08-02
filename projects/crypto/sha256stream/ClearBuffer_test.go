package crypto

import (
	"testing"
)

func TestSha256StreamClearBuffer(t *testing.T) {
	stream := Sha256Stream{}
	stream.buffer = [64]byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
		0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
	}
	stream.bitNdx = 127
	stream.byteNdx = 127

	stream.ClearBuffer(true)

	// Check if the buffer has been cleared (all elements are 0) after calling ClearBuffer
	for i := 0; i < 64; i++ {
		if stream.buffer[i] != 0x00 {
			t.Errorf("ClearBuffer() method didn't clear the buffer as expected.\n"+
				"Expected buffer[%d]: 0\n"+
				"Actual buffer[%d]: %d", i, i, stream.buffer[i])
		}
	}
	if stream.bitNdx != 0 {
		t.Fatal("bitNdx expects 0")
	}
	if stream.byteNdx != 0 {
		t.Fatal("byteNdx expects 0")
	}
	// Check if bitNdx and byteNdx have been reset to 0 after calling ClearBuffer
	if stream.bitNdx != 0 {
		t.Errorf("ClearBuffer() method didn't reset bitNdx as expected.\n"+
			"Expected bitNdx: 0\n"+
			"Actual bitNdx: %d", stream.bitNdx)
	}
	if stream.byteNdx != 0 {
		t.Errorf("ClearBuffer() method didn't reset byteNdx as expected.\n"+
			"Expected byteNdx: 0\n"+
			"Actual byteNdx: %d", stream.byteNdx)
	}
}
