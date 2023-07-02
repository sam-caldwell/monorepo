package counters

import (
	ansi "github.com/sam-caldwell/go/v2/projects/ansi/Tester"
	"testing"
)

func TestConditional_DecrementIf(t *testing.T) {
	test := ansi.Test(t)
	defer test.Stats()
	var count Conditional
	if count.value != 0 {
		t.Fail()
	}
	i := 0
	if v, _ := count.IncrementIf(false); v != -1 {
		test.Fatal("expect -1")
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.IncrementIf(false); v != -1 {
		test.Fatal("expect -1")
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.IncrementIf(true); v != 0 {
		test.Fatal("expect 0")
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.IncrementIf(true); v != 1 {
		test.Fatal("expect 1")
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.IncrementIf(true); v != 2 {
		test.Fatal("expect 2")
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.IncrementIf(false); v != -1 {
		test.Fatal("expect -1")
	}
	test.Passf("step: %d\n", i)
	i++
	if count.value != 3 {
		test.Fatal("final state expects 3")
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.DecrementIf(false); v != -1 {
		test.Fatalf("5.expect -1 Got:%d", v)
	}
	test.Passf("step: %d\n", i)
	i++
	if count.value != 3 {
		test.Fatal("final state expects 3")
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.DecrementIf(false); v != -1 {
		test.Fatalf("6.expect -1 Got:%d", v)
	}
	test.Passf("step: %d\n", i)
	i++
	if count.value != 3 {
		test.Fatal("final state expects 3")
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.DecrementIf(true); v != 3 {
		test.Fatalf("7.expect 3 Got:%d", v)
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.DecrementIf(true); v != 2 {
		test.Fatalf("8.expect 2 Got:%d", v)
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.DecrementIf(true); v != 1 {
		test.Fatalf("9.expect 1 Got:%d", v)
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.DecrementIf(true); v != 0 {
		test.Fatalf("10.expect 0 Got:%d", v)
	}
	test.Passf("step: %d\n", i)
	i++
	if v, _ := count.DecrementIf(true); v != -1 {
		test.Fatalf("10.expect -1 Got:%d", v)
	}
	test.Passf("step: %d\n", i)
	i++
}
