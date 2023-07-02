package counters

import (
	ansi "github.com/sam-caldwell/go/v2/projects/ansi/Tester"
	"testing"
)

func TestConditional_struct(t *testing.T) {
	test := ansi.Test(t)
	defer test.Stats()
	var count Conditional
	if count.value != 0 {
		test.Fail("initial state")
	}
	test.Pass("initial state")
}
