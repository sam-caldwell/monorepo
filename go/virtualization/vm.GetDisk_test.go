package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk"
	"testing"
)

func TestVM_GetDisk(t *testing.T) {
	var vm VM
	if d := vm.GetDisk(0); (d.GetName() != "") && (d.GetSize() != 0) {
		t.Fatalf("initial disk name state")
	}
	disks := make([]vmDisk.DiskDescriptor, 2)
	_ = disks[0].Name("disk0")
	_ = disks[0].Size(10)
	_ = disks[1].Name("disk1")
	_ = disks[1].Size(20)
	vm.disks = disks

	if d := vm.GetDisk(0); (d.GetName() != "disk0") && (d.GetSize() != 10) {
		t.Fatalf("disk0 mismatch")
	}
	if d := vm.GetDisk(1); (d.GetName() != "disk1") && (d.GetSize() != 20) {
		t.Fatalf("disk1 mismatch")
	}
}
