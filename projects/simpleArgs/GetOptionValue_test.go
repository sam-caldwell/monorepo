package simpleArgs

import (
	"fmt"
	"os"
	"testing"
)

func TestGetOptionValue(t *testing.T) {
	os.Args = append(os.Args, "--test0", "--test1", "value1", "--test2", "value", "--test3", "--test4")
	if value, err := GetOptionValue("--test1"); err != nil {
		t.Fatal("test1 unexpected error")
	} else {
		if value != "value1" {
			t.Fatal("test1 mismatch")
		}
	}
	if _, err := GetOptionValue("--test3"); err == nil {
		t.Fatal("Expected \"option has no value\" error not found")

	} else {
		expectedError := fmt.Sprintf("insufficient arguments for %s", "--test3")
		if err.Error() != expectedError {
			t.Fatalf("error message not right\nActual:%s\nExpected:%s", err.Error(), expectedError)
		}
	}
}
