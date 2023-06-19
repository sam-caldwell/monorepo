package simple

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"testing"
)

func TestSet_Add(t *testing.T) {
	func() {
		//Simple test
		var set Set
		if err := set.Add(1); err != nil {
			t.Fatalf("simple test: %v", err)
		}
		if err := set.Add(2); err != nil {
			t.Fatalf("simple test: %v", err)
		}
	}()
	func() {
		//Complex test
		set := Set{}

		complexTest := func(value any, expectedError error) {
			// Add the first item (type int)
			t.Logf("Value: %v [expected:%v]", value, expectedError)
			if err := set.Add(value); (err != nil) && (err.Error() != expectedError.Error()) {
				t.Fatalf("Expected error mismatch.\n"+
					"\t   value: '%v'\n"+
					"\tExpected: '%v'\n"+
					"\t     Got: '%v'\n",
					value, expectedError, err)
			}
		}

		for i := 42; i > -42; i-- {
			complexTest(42, nil)
		}

		complexTest(true, fmt.Errorf(exit.ErrTypeMismatch))
		complexTest(false, fmt.Errorf(exit.ErrTypeMismatch))
		complexTest(3.1415, fmt.Errorf(exit.ErrTypeMismatch))
		complexTest(0.0, fmt.Errorf(exit.ErrTypeMismatch))
		complexTest("", fmt.Errorf(exit.ErrTypeMismatch))
		complexTest("badString", fmt.Errorf(exit.ErrTypeMismatch))
		complexTest(nil, nil)
	}()
}
