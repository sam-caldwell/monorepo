package counters

import (
	"github.com/sam-caldwell/monorepo/go/testing/testmessage"
	"testing"
)

func TestSimpleCounter_struct(t *testing.T) {
	test := testmessage.Test(t)
	defer test.Stats()
	var count Simple
	if count.value != 0 {
		test.Fail("initial state")
	}
	test.Pass("initial state")
}
