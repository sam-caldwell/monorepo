package simple

import "testing"

func TestSet_Struct(t *testing.T) {
	var set Set
	if set.data != nil {
		t.Fatal("Initial state should be nil")
	}
}
