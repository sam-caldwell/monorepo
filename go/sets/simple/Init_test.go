package simple

import "testing"

func TestSet_Init(t *testing.T) {
	var set Set
	if set.data != nil {
		t.Fatal("expect nil data initially")
	}
	set.Init()
	if set.data == nil {
		t.Fatal("expect non-nil data after .Init()")
	}
}
