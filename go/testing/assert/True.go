package assert

import "testing"

// True - Print message and terminate test if condition is not true
func True(t *testing.T, condition bool, message string) {
	if !condition {
		t.Fatal(message)
	}
}
