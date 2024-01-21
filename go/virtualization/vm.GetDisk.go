package virtualization

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk"
)

// GetDisk - Return the given disk descriptor, identified by the index
func (vm *VM) GetDisk(index uint) vmDisk.DiskDescriptor {

	if index >= uint(len(vm.disks)) {

		vm.errors.Push(fmt.Errorf("index out of range when calling VM.GetDisks()"))

		return vmDisk.DiskDescriptor{}

	}

	return vm.disks[index]

}
