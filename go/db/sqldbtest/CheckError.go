package sqldbtest

import "testing"

// CheckError - Just make some candy to make error handling easier.
func CheckError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
