package file

import "testing"

func TestReader_Struct(t *testing.T) {
	var reader Reader
	if reader.h != nil {
		t.Fail()
	}
	if reader.buffer != nil {
		t.Fail()
	}
	if reader.filePosition != 0 {
		t.Fail()
	}
	if reader.bufferPosition != 0 {
		t.Fail()
	}
	if int8(reader.bufferBitPos) != 0 {
		t.Fail()
	}
}
