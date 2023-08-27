package convert

import (
	"reflect"
	"testing"
)

func TestPutUint64Hex(t *testing.T) {
	var v uint64 = 0x1234567890ABCDEF
	expected := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'A', 'B', 'C', 'D', 'E', 'F'}

	b := make([]byte, 16)
	PutUint64Hex(b, v)

	if !reflect.DeepEqual(b, expected) {
		t.Errorf("Unexpected result. Expected: %v, Got: %v", expected, b)
	}
}
