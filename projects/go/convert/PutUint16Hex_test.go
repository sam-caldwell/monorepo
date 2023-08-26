package convert

import (
	"reflect"
	"testing"
)

func TestPutUint16Hex(t *testing.T) {
	var v uint16 = 0xABCD
	expected := []byte{'A', 'B', 'C', 'D'}

	b := make([]byte, 4)
	PutUint16Hex(b, v)

	if !reflect.DeepEqual(b, expected) {
		t.Errorf("Unexpected result. Expected: %v, Got: %v", expected, b)
	}
}
