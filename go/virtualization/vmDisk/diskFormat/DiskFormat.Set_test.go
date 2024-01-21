package diskFormat

import "testing"

func TestDiskFormat_Set(t *testing.T) {
	var disk DiskFormat
	if err := disk.Set("ext4"); err != nil {
		t.Fatalf("unexpected error")
	}
	if err := disk.Set("swap"); err != nil {
		t.Fatalf("unexpected error")
	}
	if err := disk.Set("undefined"); err == nil {
		t.Fatalf("expected error not found")
	}
	if err := disk.Set("bad"); err == nil {
		t.Fatal("expected error not found")
	}
}
