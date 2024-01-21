package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/types/size"
	"testing"
)

func TestVM_Disk(t *testing.T) {
	var vm VM
	if len(vm.disks) != 0 {
		t.Fatal("expect initial state to be 0 disks")
	}
	if vm.errors.Size() != 0 {
		t.Fatal("expected zero reported errors initially")
	}
	vm.Disk("disk0", 10*size.MB).
		Disk("disk1", 20*size.GB)
	if vm.errors.Size() != 0 {
		t.Fatal("expected zero reported errors after adding disks")
	}
	if len(vm.disks) != 2 {
		t.Fatal("expected two disks")
	}
	if disk := vm.disks[0]; (disk.GetName() != "disk0") && disk.GetSize() != 10 {
		t.Fatal("expected disk0 descriptor mismatch")
	}
	if disk := vm.disks[1]; (disk.GetName() != "disk1") && disk.GetSize() != 20 {
		t.Fatal("expected disk0 descriptor mismatch")
	}
	vm.Disk("bad.diskName", 10)
	if vm.errors.Size() == 0 {
		t.Fatal("expect 1 error after bad disk added")
	}
}
