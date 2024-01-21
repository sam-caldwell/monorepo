package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/types/size"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk"
)

func (vm *VM) Disk(name string, sz size.Memory) *VM {
	var disk vmDisk.DiskDescriptor
	if err := disk.Name(name); err != nil {
		vm.errors.Push(err)
	}
	if err := disk.Size(sz * size.MB); err != nil {
		vm.errors.Push(err)
	}
	vm.disks = append(vm.disks, disk)
	return vm
}
