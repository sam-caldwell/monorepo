package counters

import (
	"github.com/sam-caldwell/go/v2/projects/go/ansi/Tester"
	"testing"
)

func TestSimpleCounter_struct(t *testing.T) {
	test := ansi.Test(t)
	defer test.Stats()
	var count Simple
	if count.value != 0 {
		test.Fail("initial state")
	}
	test.Pass("initial state")
}
