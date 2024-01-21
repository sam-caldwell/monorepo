package vmDisk

import "github.com/sam-caldwell/monorepo/go/virtualization/vmDisk/diskFormat"

// GetFormat - return the file format / partition type to be used for the given disk
func (disk *DiskDescriptor) GetFormat() diskFormat.DiskFormat {
	return disk.format
}
