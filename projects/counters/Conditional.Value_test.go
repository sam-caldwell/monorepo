package counters

import "testing"

func TestConditional_Value(t *testing.T) {
	var count Conditional
	if count.Value() != 0 {
		t.Fail()
	}
	if v, _ := count.Increment(); v != 0 {
		t.Fail()
	}
	if count.Value() != 1 {
		t.Fail()
	}
}
