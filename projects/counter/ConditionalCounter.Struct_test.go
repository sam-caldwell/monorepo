package counters

import "testing"

func TestConditionalCounter_struct(t *testing.T) {
	var count ConditionalCounter
	if count.value != 0 {
		t.Fail()
	}
}
