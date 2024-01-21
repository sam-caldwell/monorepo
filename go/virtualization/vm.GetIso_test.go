package virtualization

import "testing"

func TestVM_GetIso(t *testing.T) {
	var vm VM
	if vm.iso != "" {
		t.Fatal("expect vm.iso empty initially")
	}
	vm.iso = "testValue"
	if vm.iso != "testValue" {
		t.Fatal("expect vm.iso mismatch")
	}
	if iso := vm.GetIso(); iso != "testValue" {
		t.Fatal("expect VM.GetIso() mismatch")
	}
}
