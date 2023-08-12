package file

import "testing"

func TestBinaryIo_Struct(t *testing.T) {
	var io BinaryIo
	if io.handle != nil {
		t.Fail()
	}
}
