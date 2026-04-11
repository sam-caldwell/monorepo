package alphaNumericIdentifier

import (
	"testing"
)

// TestAlphaNumericIdentifier - verify the basic methods and structure of the Identifier type
func TestAlphaNumericIdentifier(t *testing.T) {
	const initialState = "initial state"
	var name Identifier
	name = Identifier(initialState)
	if string(name) != initialState {
		t.Fatal("failed to set initial state")
	}
	if err := name.Set("vmName"); err != nil {
		t.Fatal("expected no error")
	}
	if n := name.Get(); n != "vmName" {
		t.Fatalf("get mismatch.  Got %s", n)
	}
}
