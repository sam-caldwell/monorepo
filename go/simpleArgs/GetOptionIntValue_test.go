package simpleArgs

import (
	"os"
	"testing"
)

func TestGetOptionIntValue(t *testing.T) {
	os.Args = append(os.Args, "--optA", "1", "--optB", "foo", "--optC", "--optD")
	if len(os.Args) != 8 {
		t.Fatal("Test setup is wrong")
	}

	// Preflight check
	if value, err := GetOptionValue("--optA"); err != nil {
		t.Fatal("Error getting --optA in pre-check")
	} else if value != "1" {
		t.Fatal("Wrong value for --optA in pre-check")
	}
	if value, err := GetOptionValue("--optB"); err != nil {
		t.Fatal("Error getting --optB in pre-check")
	} else if value != "foo" {
		t.Fatal("Wrong value for --optB in pre-check")
	}
	if _, err := GetOptionValue("--optC"); err == nil {
		t.Fatal("Expected error on --optC in pre-check")
	}
	if _, err := GetOptionValue("--optD"); err == nil {
		t.Fatal("Expected error on --optD in pre-check")
	}

	// Happy Path: Option A is an integer.
	// 		* If required and provided, we are okay.
	//		* if not required and provided, we are okay.
	if value, err := GetOptionIntValue("--optA", true); err != nil {
		t.Fatal("Failed to get --optA integer when expected")
	} else if value != 1 {
		t.Fatal("Expected integer value not found")
	}
	if value, err := GetOptionIntValue("--optA", false); err != nil {
		t.Fatal("Failed to get --optA integer when expected")
	} else if value != 1 {
		t.Fatal("Expected integer value not found")
	}

	// Sad Path: Option B is not an integer...we expect errors.
	//		* If required, we have a problem because value is not an integer
	//      * If not required, we have a problem because value is not an integer.
	if value, err := GetOptionIntValue("--optB", true); err == nil {
		t.Fatal("Expected error for --optB but got none")
	} else if value != 0 {
		t.Fatalf("Expected default value (0) on error: %v", err)
	}
	if value, err := GetOptionIntValue("--optB", false); err == nil {
		t.Fatal("Expected error for --optB but got none")
	} else if value != 0 {
		t.Fatalf("Expected default value (0) on error: %v", err)
	}

	// Sad Path: Option doesn't exist...we expect errors.
	//		* If required --optNotReal doesn't exist and an error is expected.
	//		* If not required, --optNotReal doesn't exist and an error is not expected and default value is.
	if _, err := GetOptionIntValue("--optNotReal", true); err == nil {
		t.Fatal("Expected error for --optNotReal which doesn't exist and is required.")
	} else if err.Error() != "option not found" {
		t.Fatalf("Expected 'option not found error' but got %v", err.Error())
	}

	if value, err := GetOptionIntValue("--optNotReal", false); err != nil {
		t.Fatal("Unexpected error where --optNotReal is not required.")
	} else if value != 0 {
		t.Fatalf("Expected default value (0) on error: %v", err)
	}

	//
}
