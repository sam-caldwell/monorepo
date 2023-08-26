package valid

import (
	"fmt"
	"testing"
)

func TestIsLongArg(t *testing.T) {
	data := map[string]error{
		"--arg":            nil,
		"--arg_underscore": fmt.Errorf(errExpectedLongArg, "--arg_underscore"),
		"--arg-hyphen":     fmt.Errorf(errExpectedLongArg, "--arg-hyphen"),
		"--argMixedCase":   fmt.Errorf(errExpectedLongArg, "--argMixedCase"),
		"":                 fmt.Errorf(errExpectedLongArg, ""),
	}

	for argument, expectedErr := range data {
		if err := IsLongArg(&argument); err != nil {
			if expectedErr == nil {
				t.Fatalf("expected error on %s", argument)
			}
			if expectedErr.Error() != err.Error() {
				t.Fatalf("error mismatch (%s): %v", argument, err)
			}
		}
	}
}

func TestIsLongArg_NullInput(t *testing.T) {
	err := IsLongArg(nil)
	if err == nil {
		t.Fatal("Expected error on nil input")
	} else if err.Error() != errEmptyOrNilObject {
		t.Fatal("error string mismatch on nil input")
	}
}
