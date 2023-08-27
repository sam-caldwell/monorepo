package valid

import (
	"fmt"
	"testing"
)

func TestIsShortArg(t *testing.T) {
	data := map[string]error{
		"-a":               nil,
		"-1":               nil,
		"--arg_underscore": fmt.Errorf(errExpectedShortArg, "--arg_underscore"),
		"-arg-hyphen":      fmt.Errorf(errExpectedShortArg, "-arg-hyphen"),
		"-argMixedCase":    fmt.Errorf(errExpectedShortArg, "-argMixedCase"),
		"":                 fmt.Errorf(errExpectedShortArg, ""),
	}

	for argument, expectedErr := range data {
		if err := IsShortArg(&argument); err != nil {
			if expectedErr == nil {
				t.Fatalf("expected error on \"%s\"", argument)
			}
			if expectedErr.Error() != err.Error() {
				t.Fatalf("error mismatch (%s): \"%v\"", argument, err)
			}
		}
	}
}

func TestIsShortArg_NullInput(t *testing.T) {
	err := IsShortArg(nil)
	if err == nil {
		t.Fatal("Expected error on nil input")
	} else if err.Error() != errEmptyOrNilObject {
		t.Fatal("error string mismatch on nil input")
	}
}
