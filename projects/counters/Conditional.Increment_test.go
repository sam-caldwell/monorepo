package counters

import "testing"

func TestConditional_Increment(t *testing.T) {
	var count Conditional
	if count.value != 0 {
		t.Fail()
	}
	if v, _ := count.Increment(); v != 0 {
		t.Fatal("expect 0")
	}
	if v, _ := count.Increment(); v != 1 {
		t.Fatal("expect 1")
	}
	if v, _ := count.Increment(); v != 2 {
		t.Fatal("expect 2")
	}
}
