package counters

import "testing"

func TestSimpleCounter_struct(t *testing.T) {
	var count Simple
	if count.value != 0 {
		t.Fail()
	}
}
