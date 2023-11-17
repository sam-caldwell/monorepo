package vartype

import (
	"reflect"
	"testing"
)

func TestGeneric_Type(t *testing.T) {
	var o Generic

	o.data = 1
	if o.Type() != reflect.Int {
		t.Fail()
	}
}
