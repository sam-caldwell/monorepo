package counters

import "testing"

func TestConditional_Decrement(t *testing.T) {
	var count Conditional
	if count.value != 0 {
		t.Fail()
	}
	if v, _ := count.Increment(); v != 0 {
		t.Fatalf("1.expect 0 Got:%d", v)
	}
	if v, _ := count.Increment(); v != 1 {
		t.Fatalf("2.expect 1 Got:%d", v)
	}
	if v, _ := count.Increment(); v != 2 {
		t.Fatalf("3.expect 2 Got:%d", v)
	}
	if v, _ := count.Increment(); v != 3 {
		t.Fatalf("4.expect 3 Got:%d", v)
	}
	if v, _ := count.Decrement(); v != 4 {
		t.Fatalf("5.expect 4 Got:%d", v)
	}
	if v, _ := count.Decrement(); v != 3 {
		t.Fatalf("6.expect 3 Got:%d", v)
	}
	if v, _ := count.Decrement(); v != 2 {
		t.Fatalf("7.expect 2 Got:%d", v)
	}
	if v, _ := count.Decrement(); v != 1 {
		t.Fatalf("8.expect 1 Got:%d", v)
	}
	if v, _ := count.Decrement(); v != 0 {
		t.Fatalf("9.expect 0 Got:%d", v)
	}
	if v, _ := count.Decrement(); v != 0 {
		t.Fatalf("10.expect 0 Got:%d", v)
	}
}
