package simple

import "testing"

func TestSetMap_type(t *testing.T) {
	var s SetMap
	if s != nil {
		t.Fail()
	}
}
