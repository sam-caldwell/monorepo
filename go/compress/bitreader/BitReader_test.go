package bitreader

import (
	"testing"
)

func TestBitReader_Struct(t *testing.T) {
	var br BitReader
	if br.reader != nil {
		t.Error("BitReader should have a nil reader")
	}
	if br.buffer != uint64(0) {
		t.Error("BitReader should have a 0 buffer value")
	}

	if br.numberBitsInBuffer != uint64(0) {
		t.Error("BitReader should have a 0 numberBitsInBuffer value")
	}

	if br.err != nil {
		t.Error("BitReader should have a nil buffer")
	}
}
