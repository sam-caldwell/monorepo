package vmDisk

// GetName - return the name of the disk (used for the disk image filename)
func (disk *DiskDescriptor) GetName() string {
	return disk.name.Get()
}
