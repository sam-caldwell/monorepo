package diskFormat

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"testing"
)

func TestDiskFormat_String(t *testing.T) {

	var disk DiskFormat

	if disk.String() != words.Undefined {
		t.Fatal("expected 'undefined' but got something else")
	}

	if disk = Ext4; disk.String() != "ext4" {
		t.Fatal("expected ext4")
	}

	if disk = Swap; disk.String() != "swap" {
		t.Fatal("expected swap")
	}

}
