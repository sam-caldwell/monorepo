package simple

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/projects/v2/exit/errors"
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
			//t.Logf("Value: %v [expected:%v]", value, expectedError)
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

		complexTest(true, fmt.Errorf(errors.TypeMismatch))
		complexTest(false, fmt.Errorf(errors.TypeMismatch))
		complexTest(3.1415, fmt.Errorf(errors.TypeMismatch))
		complexTest(0.0, fmt.Errorf(errors.TypeMismatch))
		complexTest("", fmt.Errorf(errors.TypeMismatch))
		complexTest("badString", fmt.Errorf(errors.TypeMismatch))
		complexTest(nil, nil)
	}()
}
