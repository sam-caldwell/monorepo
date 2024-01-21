package virtualization

import "testing"

func TestVM_GetName(t *testing.T) {
	var vm VM
	if n := vm.name.Get(); n != "" {
		t.Fatalf("name initial state is wrong")
	}
	if err := vm.name.Set("myTestName01"); err != nil {
		t.Fatalf("error setting test value")
	}
	if n := vm.GetName(); n != "myTestName01" {
		t.Fatalf("expected value mismatch")
	}
}
