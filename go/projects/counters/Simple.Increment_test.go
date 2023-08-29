package counters

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/ansi/Tester"
	"testing"
)

func TestSimpleCounter_Increment(t *testing.T) {
	test := ansi.Test(t)
	defer test.Stats()
	var count Simple
	if count.value != 0 {
		test.Fail("initial state")
	}
	if count.Increment() != 0 {
		test.Fatal("expect 0")
	}
	if count.Increment() != 1 {
		test.Fatal("expect 1")
	}
	if count.Increment() != 2 {
		test.Fatal("expect 2")
	}
}
