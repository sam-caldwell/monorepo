package counters

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/ansi/Tester"
	"testing"
)

func TestSimpleCounter_Decrement(t *testing.T) {
	test := ansi.Test(t)
	defer test.Stats()
	var count Simple
	if count.value != 0 {
		test.Fail("initial state")
	}
	if v := count.Increment(); v != 0 {
		test.Fatalf("1.expect 0 Got:%d", v)
	}
	if v := count.Increment(); v != 1 {
		test.Fatalf("2.expect 1 Got:%d", v)
	}
	if v := count.Increment(); v != 2 {
		test.Fatalf("3.expect 2 Got:%d", v)
	}
	if v := count.Increment(); v != 3 {
		test.Fatalf("4.expect 3 Got:%d", v)
	}
	if v := count.Decrement(); v != 4 {
		test.Fatalf("5.expect 4 Got:%d", v)
	}
	if v := count.Decrement(); v != 3 {
		test.Fatalf("6.expect 3 Got:%d", v)
	}
	if v := count.Decrement(); v != 2 {
		test.Fatalf("7.expect 2 Got:%d", v)
	}
	if v := count.Decrement(); v != 1 {
		test.Fatalf("8.expect 1 Got:%d", v)
	}
	if v := count.Decrement(); v != 0 {
		test.Fatalf("9.expect 0 Got:%d", v)
	}
	if v := count.Decrement(); v != -1 {
		test.Fatalf("10.expect -1 Got:%d", v)
	}
}
