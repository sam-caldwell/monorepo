package vmDisk

import "github.com/sam-caldwell/monorepo/go/virtualization/vmDisk/diskFormat"

// Format - specify the file format / partition type used for the disk image.
func (disk *DiskDescriptor) Format(format diskFormat.DiskFormat) error {
	disk.format = format
	return nil
}
