package counters

import "testing"

func TestConditional_IncrementIf(t *testing.T) {
	var count Conditional
	if count.value != 0 {
		t.Fail()
	}
	if v, _ := count.IncrementIf(false); v != -1 {
		t.Fatal("expect -1")
	}
	if v, _ := count.IncrementIf(false); v != -1 {
		t.Fatal("expect -1")
	}
	if v, _ := count.IncrementIf(true); v != 0 {
		t.Fatal("expect 0")
	}
	if v, _ := count.IncrementIf(true); v != 1 {
		t.Fatal("expect 1")
	}
	if v, _ := count.IncrementIf(true); v != 2 {
		t.Fatal("expect 2")
	}
	if v, _ := count.IncrementIf(false); v != -1 {
		t.Fatal("expect -1")
	}
	if count.value != 3 {
		t.Fatal("final state expects 3")
	}
}
