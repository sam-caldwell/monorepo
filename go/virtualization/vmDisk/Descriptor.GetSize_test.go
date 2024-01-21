package vmDisk

import (
	"github.com/sam-caldwell/monorepo/go/types/size"
	"testing"
)

func TestDiskDescriptor_GetSize(t *testing.T) {
	var disk DiskDescriptor
	disk.size = 10 * size.GB
	if sz := disk.GetSize(); sz != uint(10*size.GB) {
		t.Fatal("expected GetSize() mismatch")
	}
}
