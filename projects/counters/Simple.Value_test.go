package counters

import "testing"

func TestSimpleCounter_Value(t *testing.T) {
	var count SimpleCounter
	if count.Value() != 0 {
		t.Fail()
	}
	if count.Increment() != 0 {
		t.Fail()
	}
	if count.Value() != 1 {
		t.Fail()
	}
}
