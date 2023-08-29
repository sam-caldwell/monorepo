package counters

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/ansi/Tester"
	"testing"
)

func TestSimpleCounter_Value(t *testing.T) {
	test := ansi.Test(t)
	defer test.Stats()
	var count Simple
	if count.Value() != 0 {
		test.Fail("initial state")
	}
	test.Pass("initial state")
	if count.Increment() != 0 {
		test.Fail("expect 0")
	}
	test.Pass("expect 0")
	if count.Value() != 1 {
		test.Fail("expect 1")
	}
	test.Pass("expect 1")
}
