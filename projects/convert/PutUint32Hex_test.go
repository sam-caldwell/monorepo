package convert

import (
	"reflect"
	"testing"
)

func TestPutUint32Hex(t *testing.T) {
	var v uint32 = 0x12345678
	expected := []byte{'1', '2', '3', '4', '5', '6', '7', '8'}

	b := make([]byte, 8)
	PutUint32Hex(b, v)

	if !reflect.DeepEqual(b, expected) {
		t.Errorf("Unexpected result. Expected: %v, Got: %v", expected, b)
	}
}
