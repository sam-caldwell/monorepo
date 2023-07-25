package assert

import "testing"

// False - Print message and terminate test if condition is not false
func False(t *testing.T, condition bool, message string) {
	if condition {
		t.Fatal(message)
	}
}
