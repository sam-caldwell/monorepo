package ordered

import "testing"

func TestSet_Struct(t *testing.T) {
	var s Set
	if s.data != nil {
		t.Fail()
	}
	if s.typ != nil {
		t.Fail()
	}
}
