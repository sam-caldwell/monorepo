package counters

import "testing"

func TestConditional_struct(t *testing.T) {
	var count Conditional
	if count.value != 0 {
		t.Fail()
	}
}
