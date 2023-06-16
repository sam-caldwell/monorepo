package counters

import "testing"

func TestConditionalCounter_Value(t *testing.T) {
	var count ConditionalCounter
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
