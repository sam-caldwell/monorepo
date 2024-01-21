package vmDisk

import (
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk/diskFormat"
	"testing"
)

func TestDiskDescriptor_Format(t *testing.T) {

	var disk DiskDescriptor

	if err := disk.Format(diskFormat.Ext4); err != nil {
		t.Fatal("expect no error")
	}

	if err := disk.Format(diskFormat.Swap); err != nil {
		t.Fatal("expect no error")
	}

}
