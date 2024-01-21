package vmDisk

import "github.com/sam-caldwell/monorepo/go/types/alphaNumericIdentifier"

type DiskDescriptor struct {
	name alphaNumericIdentifier.Identifier
	size uint //size in MB
	//ToDo: Add other things...?
}

func (disk *DiskDescriptor) GetName() string {
	return disk.name.Get()
}

func (disk *DiskDescriptor) GetSize() uint {
	return disk.size
}

func (disk *DiskDescriptor) Name(n string) error {
	return disk.name.Set(n)
}

func (disk *DiskDescriptor) Size(sz uint) error {
	disk.size = sz
	return nil
}
