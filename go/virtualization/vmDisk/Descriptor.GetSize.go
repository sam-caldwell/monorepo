package vmDisk

// GetSize - Return the disk size (in bytes)
func (disk *DiskDescriptor) GetSize() uint {
	return uint(disk.size)
}
