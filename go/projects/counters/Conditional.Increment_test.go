package counters

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/testing/testmessage"
	"testing"
)

func TestConditional_Increment(t *testing.T) {
	test := testmessage.Test(t)
	defer test.Stats()
	var count Conditional
	if count.value != 0 {
		test.Fail("initial state")
	}
	test.Pass("initial state")
	if v, _ := count.Increment(); v != 0 {
		test.Fatal("expect 0")
	}
	test.Pass("expect 0")
	if v, _ := count.Increment(); v != 1 {
		test.Fatal("expect 1")
	}
	test.Pass("expect 1")
	if v, _ := count.Increment(); v != 2 {
		test.Fatal("expect 2")
	}
	test.Pass("expect 2")
}
