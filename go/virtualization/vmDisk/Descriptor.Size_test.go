package vmDisk

import (
	"github.com/sam-caldwell/monorepo/go/types/size"
	"testing"
)

func TestDiskDescriptor_Size(t *testing.T) {

	var disk DiskDescriptor

	if disk.size != 0 {
		t.Fatal("expect size 0 initially")
	}

	if err := disk.Size(10 * size.GB); err != nil {
		t.Fatal(err)
	}

	if disk.size != 10*size.GB {
		t.Fatal("size mismatch")
	}

}
