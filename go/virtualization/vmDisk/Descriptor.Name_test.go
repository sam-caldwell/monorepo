package vmDisk

import "testing"

func TestDiskDescriptor_Name(t *testing.T) {

	var disk DiskDescriptor

	if d := disk.name.Get(); d != "" {
		t.Fatal("expect empty state")
	}

	if err := disk.Name("testName"); err != nil {
		t.Fatal(err)
	}

	if d := disk.name.Get(); d != "testName" {
		t.Fatal("expect name mismatch")
	}

}
