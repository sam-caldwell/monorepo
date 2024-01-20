package vmname

import (
	"testing"
)

// TestVmName - verify the basic methods and structure of the VmName type
func TestVmName(t *testing.T) {
	const initialState = "initial state"
	var name VmName
	name = VmName(initialState)
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
