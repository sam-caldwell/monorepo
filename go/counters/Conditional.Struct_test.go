package counters

import (
	"github.com/sam-caldwell/monorepo/go/testing/testmessage"
	"testing"
)

func TestConditional_struct(t *testing.T) {
	test := testmessage.Test(t)
	defer test.Stats()
	var count Conditional
	if count.value != 0 {
		test.Fail("initial state")
	}
	test.Pass("initial state")
}
