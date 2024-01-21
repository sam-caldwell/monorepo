package vmDisk

import "testing"

func TestDiskDescriptor_GetName(t *testing.T) {
	var disk DiskDescriptor

	if d := disk.GetName(); d != "" {
		t.Fatal("expect empty name")
	}

	disk.name = "testName"

	if d := disk.GetName(); d != "testName" {
		t.Fatal("expect testName")
	}

}
