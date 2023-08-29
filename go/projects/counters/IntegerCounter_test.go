package counters

import (
	"github.com/sam-caldwell/monorepo/go/projects/v2/ansi/Tester"
	"testing"
)

func TestIntegerCounter_incrementing(t *testing.T) {
	test := ansi.Test(t)
	defer test.Stats()

	count := IntegerCounter(0, 1)

	testFunction := func(n int) {
		if count() != n {
			test.Fatalf("testing %d\n", n)
		}
		//test.Passf("testing %d\n", n)
	}

	for i := 1; i <= 100; i++ {
		testFunction(i)
	}
}

func TestIntegerCounter_decrementing(t *testing.T) {
	test := ansi.Test(t)
	defer test.Stats()

	count := IntegerCounter(10, -1)

	testFunction := func(n int) {
		if v := count(); v != n-1 {
			test.Fatalf("testing %d Got: %d\n", n, v)
		}
		//test.Passf("testing %d\n", n)
	}

	for i := 10; i > 0; i-- {
		testFunction(i)
	}

}
