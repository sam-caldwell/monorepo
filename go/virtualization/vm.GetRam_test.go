package virtualization

import "testing"

func TestVM_GetRam(t *testing.T) {
	var vm VM
	if r := vm.GetRam(); r != 0 {
		t.Fatalf("initial state mismatch")
	}
	vm.ram = 1337
	if r := vm.GetRam(); r != 1337 {
		t.Fatalf("set state mismatch")
	}
}
