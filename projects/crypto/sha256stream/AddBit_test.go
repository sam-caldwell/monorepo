package crypto

import (
	"testing"
)

func TestSha256StreamAddBit_fullBuffer(t *testing.T) {
	// Create a new Sha256Stream instance
	stream := Sha256Stream{}

	// Add 64 bits to the buffer, which should trigger processing the message block
	for i := 0; i < 64; i++ {
		stream.AddBit(true)
	}

	// Calculate the expected bitNdx after adding 64 bits
	expectedBitNdx := int8(64)

	// Check if the buffer has been processed (i.e., byteNdx becomes 0) after adding 64 bits
	if stream.byteNdx != int8(0) {
		t.Errorf("AddBit() method didn't process the buffer into the hash state as expected.\n"+
			"Expected byteNdx: %d\n"+
			"Actual byteNdx: %d", int8(len(stream.buffer)), stream.byteNdx)
	}

	// Check if the bitNdx has been updated correctly after adding 64 bits
	if stream.bitNdx != expectedBitNdx {
		t.Errorf("AddBit() method didn't update the bitNdx as expected.\n"+
			"Expected bitNdx: %d\n"+
			"Actual bitNdx: %d", expectedBitNdx, stream.bitNdx)
	}
}

func TestSha256StreamAddBit_HalfBuffer(t *testing.T) {
	// Create a new Sha256Stream instance
	stream := Sha256Stream{}

	// Add 64 bits to the buffer, which should trigger processing the message block
	for i := 0; i < 32; i++ {
		stream.AddBit(true)
	}

	// Calculate the expected bitNdx after adding 64 bits
	expectedBitNdx := int8(32)

	// Check if the buffer has been processed (i.e., byteNdx becomes 0) after adding 64 bits
	if stream.byteNdx != int8(0) {
		t.Errorf("AddBit() method didn't process the buffer into the hash state as expected.\n"+
			"Expected byteNdx: %d\n"+
			"Actual byteNdx: %d", int8(len(stream.buffer)), stream.byteNdx)
	}

	// Check if the bitNdx has been updated correctly after adding 64 bits
	if stream.bitNdx != expectedBitNdx {
		t.Errorf("AddBit() method didn't update the bitNdx as expected.\n"+
			"Expected bitNdx: %d\n"+
			"Actual bitNdx: %d", expectedBitNdx, stream.bitNdx)
	}
}
