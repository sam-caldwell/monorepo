package ordered

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"reflect"
	"testing"
)

func TestOrderedSet_Add_Handles_Arbitrary_Type_Inputs(t *testing.T) {
	/*
	 * Test assumptions...
	 *  1. We can use many different data types in a set.
	 *  2. If we pass an arbitrary type object into the set, it will be added without issue.
	 *  3. Once we have added a value, the set should now be of the type of the value added.
	 */

	// Setup a test function...
	simpleTest := func(data any, expectedError error) {
		// Create a new set
		set := Set{}
		// Verify empty set
		if len(set.data) != 0 {
			t.Fatal("set should be empty initially")
		}
		// Add data to the set
		if err := set.Add(data); (err != nil) && ((expectedError == nil) || (err.Error() != expectedError.Error())) {
			t.Fatalf("Expected '%v', got: '%v'", expectedError, err)
		}
		// Make sure the data was added to the set.
		if len(set.data) == 0 {
			t.Fatalf("exected a record to exist")
		}
		// Verify the type is now set.
		if reflect.TypeOf(set.data[0]).Kind() != reflect.TypeOf(data).Kind() {
			t.Fatalf("Type was not set when the first element was added.\n"+
				"Expected: %v\n"+
				"     Got: %v",
				reflect.TypeOf(set.data[0]).Kind().String(),
				reflect.TypeOf(data).Kind().String())
		}
		if expectedError != nil {
			t.Fatal("Expected error not encountered")
		}
	}

	//Use our test function to run against the expected inputs.
	simpleTest(1, nil)
	simpleTest("hi", nil)
	simpleTest(true, nil)
	simpleTest(false, nil)
	simpleTest(1.0, nil)
	simpleTest(-1, nil)
}

func TestOrderedSet_Add_Enforces_TypeChecking(t *testing.T) {
	/*
	 * Test assumptions...
	 *  1. We can use many different data types in a set.
	 *  2. If we pass an arbitrary type object into the set, it will be added without issue.
	 *  3. Once we have added a value, the set should now be of the type of the value added.
	 *  4. The second and subsequent records will have type-checks enforced.
	 */

	// Create our set.  We will have only one set to rule them all...
	set := Set{}

	// Setup a test function...
	testFunction := func(count int, data any, expectedError error) {
		// Add our new data to the set.
		if err := set.Add(data); err != nil {
			// We have an error.  Do we expect one?
			if expectedError == nil {
				t.Fatalf("On '%v', expected no error but got '%v'\n"+
					" Set Type: '%v'\n"+
					"Data Type: '%v'\n",
					data, err, reflect.TypeOf(set.data[0]),
					reflect.TypeOf(data))
			} else {
				if expectedError.Error() != err.Error() {
					t.Fatalf("On %v, we got an error other than what we expected.\n"+
						"Expected: %s\n"+
						"     Got: %s\n",
						data,
						expectedError.Error(),
						err.Error())
				}
				return
			}
		} // .Add() has no error...
		if expectedError != nil {
			t.Fatalf("We expected an error but got none on '%v'", data)
		}
		if len(set.data) != count {
			t.Fatalf("expected '%d' records.  Got count: '%d' on value: %v", count, set.Count(), data)
		}
	}

	// Use our test function to test various data scenarios.
	testFunction(1, 1, nil)
	testFunction(1, "hi", fmt.Errorf(errors.TypeMismatch))
	testFunction(1, true, fmt.Errorf(errors.TypeMismatch))
	testFunction(1, false, fmt.Errorf(errors.TypeMismatch))
	testFunction(1, 1.0, fmt.Errorf(errors.TypeMismatch))
	testFunction(2, -1, nil)
	//test(3, 0, nil)
}
