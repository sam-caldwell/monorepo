package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/types/size"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk/diskFormat"
	"testing"
)

func TestVM_Disk(t *testing.T) {

	t.Run("readonly protections", func(t *testing.T) {
		var vm VM
		if vm.readOnly {
			t.Fatalf("expect readonly false")
		}
		vm.readOnly = true
		if !vm.readOnly {
			t.Fatalf("expect readonly true")
		}
		if p := vm.Disk("disk0", 10*size.MB, diskFormat.Ext4); p != &vm {
			t.Fatalf("expect address of vm object")
		}
		if vm.errors.Size() == 0 {
			t.Fatalf("expect an error on readonly violation")
		}
		if err := vm.errors.Dump(); err[0].Error() != "readonly violation (disk)" {
			t.Fatalf("error message mismatch")
		}
	})
	t.Run("initial state", func(t *testing.T) {
		var vm VM
		if len(vm.disks) != 0 {
			t.Fatal("expect initial state to be 0 disks")
		}

		if vm.errors.Size() != 0 {
			t.Fatal("expected zero reported errors initially")
		}
	})
	t.Run("create two disks", func(t *testing.T) {
		var vm VM
		vm.Disk("disk0", 10*size.MB, diskFormat.Ext4).
			Disk("disk1", 20*size.GB, diskFormat.Swap)
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
	})
	t.Run("sad path: bad disk Name1", func(t *testing.T) {
		var vm VM

		vm.Disk("bad.diskName", 10, diskFormat.Ext4)

		if vm.errors.Size() == 0 {
			t.Fatal("expect 1 error after bad disk added")
		}
	})
	t.Run("sad path: bad disk Name2", func(t *testing.T) {
		var vm VM

		vm.Disk("bad-diskName", 10, diskFormat.Ext4)

		if vm.errors.Size() == 0 {
			t.Fatal("expect 1 error after bad disk added")
		}
	})
	t.Run("sad path: bad disk Name3", func(t *testing.T) {
		var vm VM

		vm.Disk("bad_diskName", 10, diskFormat.Ext4)

		if vm.errors.Size() == 0 {
			t.Fatal("expect 1 error after bad disk added")
		}
	})

}
