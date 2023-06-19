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
			t.Fatal(err)
		}
		if err := set.Add(2); err != nil {
			t.Fatal(err)
		}
	}()
	func() {
		//Complex test
		set := Set{}

		test := func(value any, expectedError error) {
			// Add the first item (type int)
			if err := set.Add(value); (err != nil) && (err.Error() != expectedError.Error()) {
				t.Fatalf("Expected error mismatch.\n"+
					"\t   value: '%v'\n"+
					"\tExpected: '%v'\n"+
					"\t     Got: '%v'\n",
					value, expectedError, err)
			}
		}

		for i := 42; i > -42; i-- {
			test(42, nil)
		}

		test(true, fmt.Errorf(exit.ErrTypeMismatch))
		test(false, fmt.Errorf(exit.ErrTypeMismatch))
		test(3.1415, fmt.Errorf(exit.ErrTypeMismatch))
		test(0.0, fmt.Errorf(exit.ErrTypeMismatch))
		test("", fmt.Errorf(exit.ErrTypeMismatch))
		test("badString", fmt.Errorf(exit.ErrTypeMismatch))
		test(nil, nil)
	}()
}
