package counters

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/ansi/Tester"
	"testing"
)

func TestConditional_Value(t *testing.T) {
	test := ansi.Test(t)
	defer test.Stats()
	var count Conditional
	if count.Value() != 0 {
		test.Fail("initial state")
	}
	if v, _ := count.Increment(); v != 0 {
		test.Fail("step 1")
	}
	test.Pass("step 1")
	if count.Value() != 1 {
		test.Fail("step 2")
	}
	test.Pass("step 2")
}
