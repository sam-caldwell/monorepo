package ordered

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"testing"
)

func TestOrderedSet_Add(t *testing.T) {

	//Test that we can use many types
	func() {
		test := func(data any, expectedError error) {
			// Create a new set
			set := Set{}
			if err := set.Add(data); (err != nil) && ((expectedError == nil) || (err.Error() != expectedError.Error())) {
				t.Errorf("Expected '%v', got: '%v'", expectedError, err)
			}
			if !set.Empty() {
				t.Errorf("exected a record to exist")
			}
		}
		test(1, nil)
		test("hi", nil)
		test(true, nil)
		test(false, nil)
		test(1.0, nil)
		test(-1, nil)
	}()

	//Test that we enforce typechecks
	func() {
		// Create a new set
		set := Set{}
		test := func(count int, data any, expectedError error) {
			if err := set.Add(data); (err != nil) && ((expectedError == nil) || (err.Error() != expectedError.Error())) {
				t.Errorf("Expected '%v', got: '%v'", expectedError, err)
			}
			if set.Count() != count {
				t.Errorf("expected %d records.  Got %d on %v", count, set.Count(), data)
			}
		}
		test(1, 1, nil)
		test(1, "hi", fmt.Errorf(exit.ErrTypeMismatch))
		test(1, true, fmt.Errorf(exit.ErrTypeMismatch))
		test(1, false, fmt.Errorf(exit.ErrTypeMismatch))
		test(1, 1.0, fmt.Errorf(exit.ErrTypeMismatch))
		test(2, -1, nil)
		test(3, 0, nil)
	}()
}
