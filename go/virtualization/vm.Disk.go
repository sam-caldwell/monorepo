package virtualization

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/types/size"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk/diskFormat"
)

func (vm *VM) Disk(name string, sz size.Memory, format diskFormat.DiskFormat) *VM {
	if vm.readOnly {
		vm.errors.Push(fmt.Errorf("readonly violation (disk)"))
	} else {
		var disk vmDisk.DiskDescriptor
		if err := disk.Name(name); err != nil {
			vm.errors.Push(err)
		}
		if err := disk.Size(sz * size.MB); err != nil {
			vm.errors.Push(err)
		}
		if err := disk.Format(format); err != nil {
			vm.errors.Push(err)
		}
		vm.disks = append(vm.disks, disk)
	}
	return vm
}
