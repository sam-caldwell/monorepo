package virtualization

import "testing"

func TestVM_GetIso(t *testing.T) {
	var vm VM
	if vm.iso != "" {
		t.Fatal("expect vm.iso empty initially")
	}
	if sz := vm.errors.Size(); sz != 0 {
		t.Fatal("error occurred initially")
	}
	if err := vm.iso.Set("testValue"); err != nil {
		t.Fatal(err)
	}
	if sz := vm.errors.Size(); sz != 0 {
		t.Fatal("error occurred when setting Iso")
	}
	if i := vm.GetIso(); i != "testValue" {
		t.Fatalf("expect VM.GetIso() mismatch. Got '%s' state: '%s'", i, vm.iso.Get())
	}
}
