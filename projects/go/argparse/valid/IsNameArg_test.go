package valid

import (
	"fmt"
	"testing"
)

func TestIsNameArg(t *testing.T) {
	data := map[string]error{
		"a":                nil,
		"thisarg":          nil,
		"thisArg":          nil,
		"this_Arg":         nil,
		"this-Arg":         nil,
		"th15Arg1":         nil,
		"1thisArg":         fmt.Errorf(errExpectedNameArg, "1thisArg"),
		"--arg_underscore": fmt.Errorf(errExpectedNameArg, "--arg_underscore"),
		"-arg-hyphen":      fmt.Errorf(errExpectedNameArg, "-arg-hyphen"),
		"-a":               fmt.Errorf(errExpectedNameArg, "-a"),
		"-argMixedCase":    fmt.Errorf(errExpectedNameArg, "-argMixedCase"),
		"":                 fmt.Errorf(errExpectedNameArg, ""),
	}

	for argument, expectedErr := range data {
		if err := IsNameArg(&argument); err != nil {
			if expectedErr == nil {
				t.Fatalf("expected error on \"%s\" %v", argument, expectedErr)
			}
			if expectedErr.Error() != err.Error() {
				t.Fatalf("error mismatch (%s): \"%v\"", argument, err)
			}
		}
	}
}

func TestIsNameArg_NullInput(t *testing.T) {
	err := IsNameArg(nil)
	if err == nil {
		t.Fatal("Expected error on nil input")
	} else if err.Error() != errEmptyOrNilObject {
		t.Fatal("error string mismatch on nil input")
	}
}
