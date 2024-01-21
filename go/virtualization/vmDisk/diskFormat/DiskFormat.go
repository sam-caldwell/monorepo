package diskFormat

type DiskFormat uint

const (
	Undefined DiskFormat = iota
	Ext4
	Swap
)
