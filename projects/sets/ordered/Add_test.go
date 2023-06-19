package ordered

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"testing"
)

func TestOrderedSet_Add(t *testing.T) {

	//Test that we can use many types
	func() {
		simpleTest := func(data any, expectedError error) {
			// Create a new set
			set := Set{}
			if err := set.Add(data); (err != nil) && ((expectedError == nil) || (err.Error() != expectedError.Error())) {
				t.Fatalf("Expected '%v', got: '%v'", expectedError, err)
			}
			if !set.Empty() {
				t.Fatalf("exected a record to exist")
			}
		}
		simpleTest(1, nil)
		simpleTest("hi", nil)
		simpleTest(true, nil)
		simpleTest(false, nil)
		simpleTest(1.0, nil)
		simpleTest(-1, nil)
	}()

	//Test that we enforce typechecks
	func() {
		// Create a new set
		set := Set{}
		test := func(count int, data any, expectedError error) {

			err := set.Add(data)
			if err == nil {
				if expectedError == nil {
					return //ok.  Keep going.
				}
				t.Fatalf("On '%v', expected '%v' no error and got none", data, expectedError)
			}
			if expectedError == nil {
				t.Fatalf("On '%v', expected no error and got '%v'", data, err)
			}
			if err.Error() != expectedError.Error() {
				t.Fatalf("Expected '%v', got: '%v' (count %d)", expectedError, err, count)
			}
			if set.Count() != count {
				t.Fatalf("expected '%d' records.  Got count: '%d' on value: %v", count, set.Count(), data)
			}
		}
		test(1, 1, nil)
		test(1, "hi", fmt.Errorf(exit.ErrTypeMismatch))
		test(1, true, fmt.Errorf(exit.ErrTypeMismatch))
		test(1, false, fmt.Errorf(exit.ErrTypeMismatch))
		test(1, 1.0, fmt.Errorf(exit.ErrTypeMismatch))
		test(2, -1, nil)
		//test(3, 0, nil)
	}()
}
