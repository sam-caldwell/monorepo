package valid

import (
	"testing"
)

func TestIsValidHelpString_NullInput(t *testing.T) {
	err := IsNameArg(nil)
	if err == nil {
		t.Fatal("Expected error on nil input")
	} else if err.Error() != errEmptyOrNilObject {
		t.Fatal("error string mismatch on nil input")
	}
}

func TestIsValidHelpString(t *testing.T) {
	data := map[string]error{
		"":            nil,
		"help string": nil,
	}

	for argument, expectedErr := range data {
		if err := IsValidHelpString(&argument); err != nil {
			if expectedErr == nil {
				t.Fatalf("expected error on \"%s\" %v", argument, expectedErr)
			}
			if expectedErr.Error() != err.Error() {
				t.Fatalf("error mismatch (%s): \"%v\"", argument, err)
			}
		}
	}
}
