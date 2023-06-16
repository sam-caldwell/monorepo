package counters

import "testing"

func TestSimpleCounter_struct(t *testing.T) {
	var count SimpleCounter
	if count.value != 0 {
		t.Fail()
	}
}
