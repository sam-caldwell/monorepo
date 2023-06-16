package vartype

import "testing"

func TestGeneric_Struct(t *testing.T) {
	var o Generic
	if o.data != nil {
		t.Fail()
	}
}
