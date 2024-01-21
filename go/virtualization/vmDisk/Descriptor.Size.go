package vmDisk

import "github.com/sam-caldwell/monorepo/go/types/size"

// Size - define the size of the given disk
func (disk *DiskDescriptor) Size(sz size.Memory) error {
	disk.size = sz
	return nil
}
