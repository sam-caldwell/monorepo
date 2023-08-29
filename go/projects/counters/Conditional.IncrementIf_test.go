package counters

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/ansi/Tester"
	"testing"
)

func TestConditional_IncrementIf(t *testing.T) {
	test := ansi.Test(t)
	defer test.Stats()
	var count Conditional
	if count.value != 0 {
		test.Fail("initial state")
	}
	test.Pass("initial state")
	i := 1
	if v, _ := count.IncrementIf(false); v != -1 {
		t.Fatal("expect -1")
	}
	test.Passf("step %d\n", i)
	i++
	if v, _ := count.IncrementIf(false); v != -1 {
		t.Fatal("expect -1")
	}
	test.Passf("step %d\n", i)
	i++
	if v, _ := count.IncrementIf(true); v != 0 {
		t.Fatal("expect 0")
	}
	test.Passf("step %d\n", i)
	i++
	if v, _ := count.IncrementIf(true); v != 1 {
		t.Fatal("expect 1")
	}
	test.Passf("step %d\n", i)
	i++
	if v, _ := count.IncrementIf(true); v != 2 {
		t.Fatal("expect 2")
	}
	test.Passf("step %d\n", i)
	i++
	if v, _ := count.IncrementIf(false); v != -1 {
		t.Fatal("expect -1")
	}
	test.Passf("step %d\n", i)
	i++
	if count.value != 3 {
		t.Fatal("final state expects 3")
	}
	test.Passf("step %d\n", i)
	i++
}
