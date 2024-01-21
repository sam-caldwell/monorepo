package virtualization

import "testing"

func TestVM_Name(t *testing.T) {
	var vm VM
	if vm.name != "" {
		t.Fatal("initial name should be empty")
	}
	if vm.errors.Size() != 0 {
		t.Fatal("expect zero errors")
	}
	if p := vm.Name("vmName"); p != &vm {
		t.Fatal("expect vm.Name() to return the address of vm")
	}
	if vm.errors.Size() != 0 {
		t.Fatal("expect zero errors")
	}
	if p := vm.Name("bad.vmName"); p != &vm {
		t.Fatal("expect vm.Name() to return the address of vm")
	}
	if vm.errors.Size() != 1 {
		t.Fatal("expect bad vm to cause vm to have one error")
	}
}
