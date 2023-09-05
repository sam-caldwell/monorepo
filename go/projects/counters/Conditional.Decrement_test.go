package counters

import (
	"github.com/sam-caldwell/monorepo/go/projects/testing/testmessage"
	"testing"
)

func TestConditional_Decrement(t *testing.T) {
	test := testmessage.Test(t)
	defer test.Stats()
	var count Conditional
	if count.value != 0 {
		test.Fail("initial state")
	}
	test.Pass("step 0: initial state")
	if v, _ := count.Increment(); v != 0 {
		test.Fatalf("1.expect 0 Got:%d", v)
	}
	test.Pass("step 1")
	if v, _ := count.Increment(); v != 1 {
		test.Fatalf("2.expect 1 Got:%d", v)
	}
	test.Pass("step 2")
	if v, _ := count.Increment(); v != 2 {
		test.Fatalf("3.expect 2 Got:%d", v)
	}
	test.Pass("step 3")
	if v, _ := count.Increment(); v != 3 {
		test.Fatalf("4.expect 3 Got:%d", v)
	}
	test.Pass("step 4")
	if v, _ := count.Decrement(); v != 4 {
		test.Fatalf("5.expect 4 Got:%d", v)
	}
	test.Pass("step 5")
	if v, _ := count.Decrement(); v != 3 {
		test.Fatalf("6.expect 3 Got:%d", v)
	}
	test.Pass("step 6")
	if v, _ := count.Decrement(); v != 2 {
		test.Fatalf("7.expect 2 Got:%d", v)
	}
	test.Pass("step 7")
	if v, _ := count.Decrement(); v != 1 {
		test.Fatalf("8.expect 1 Got:%d", v)
	}
	test.Pass("step 8")
	if v, _ := count.Decrement(); v != 0 {
		test.Fatalf("9.expect 0 Got:%d", v)
	}
	test.Pass("step 9")
	if v, _ := count.Decrement(); v != 0 {
		test.Fatalf("10.expect 0 Got:%d", v)
	}
	test.Pass("step 10")
}
