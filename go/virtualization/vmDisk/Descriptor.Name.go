package vmDisk

// Name - define the logical name of the disk (used to name the disk image file)
func (disk *DiskDescriptor) Name(n string) error {
	return disk.name.Set(n)
}
