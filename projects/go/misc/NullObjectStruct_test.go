package misc

import (
	"reflect"
	"testing"
)

func TestNullObjectStructSize(t *testing.T) {
	var o NullObjectStruct
	typ := reflect.TypeOf(o)
	if typ.Size() != 0 {
		t.Fatal("Expect NullObjectStruct to be zero-length")
	}
	if typ.Kind() != reflect.Struct {
		t.Fatal("Expect NullObjectStruct to be a struct")
	}
	if typ.NumField() != 0 {
		t.Fatal("Expect NullObjectStruct to have no fields")
	}
}
