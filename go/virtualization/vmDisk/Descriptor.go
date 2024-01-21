package vmDisk

import (
	"github.com/sam-caldwell/monorepo/go/types/alphaNumericIdentifier"
	"github.com/sam-caldwell/monorepo/go/types/size"
)

type DiskDescriptor struct {
	name alphaNumericIdentifier.Identifier
	size size.Memory
	//ToDo: Add other things...?
}

func (disk *DiskDescriptor) GetName() string {
	return disk.name.Get()
}

func (disk *DiskDescriptor) GetSize() uint {
	return uint(disk.size)
}

func (disk *DiskDescriptor) Name(n string) error {
	return disk.name.Set(n)
}

func (disk *DiskDescriptor) Size(sz size.Memory) error {
	disk.size = sz
	return nil
}
