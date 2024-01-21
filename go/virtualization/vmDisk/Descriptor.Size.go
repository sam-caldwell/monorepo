package vmDisk

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/types/size"
)

// Size - define the size of the given disk
func (disk *DiskDescriptor) Size(sz size.Memory) error {
	if uint(sz) == 0 {
		return fmt.Errorf("cannot set size to 0")
	}
	disk.size = sz
	return nil
}
