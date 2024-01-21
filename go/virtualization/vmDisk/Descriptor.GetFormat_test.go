package vmDisk

import (
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk/diskFormat"
	"testing"
)

func TestDiskDescriptor_GetFormat(t *testing.T) {

	var disk DiskDescriptor

	if d := disk.GetFormat(); d != diskFormat.Undefined {
		t.Fatal("expect undefined initially")
	}

	disk.format = diskFormat.Ext4

	if d := disk.GetFormat(); d != diskFormat.Ext4 {
		t.Fatal("expect ext4")
	}

	disk.format = diskFormat.Swap

	if d := disk.GetFormat(); d != diskFormat.Swap {
		t.Fatal("expect swap")
	}

}
