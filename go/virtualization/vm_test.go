package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/types/size"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk/diskFormat"
	"testing"
)

func TestVm(t *testing.T) {

	t.Run("Setup the vm object only", func(t *testing.T) {
		vm := NewVm().
			Name("myTestVm").
			Cpu(2).
			Ram(1024*size.MB).
			Image("{{home}}/vmImages/{{name}}").
			Iso("{{home}}/iso/ubuntu.iso").
			Disk("opsysRoot", 10*size.GB, diskFormat.Ext4).
			Disk("swap", 1*size.GB, diskFormat.Swap)
		errors := vm.GetErrors()
		for n, err := range errors {
			t.Errorf("error[%d]: %v", n, err)
		}
		if len(errors) != 0 {
			t.Fatal("has errors")
		}
	})
	//ToDo: Add Create test
	//ToDo: Add Start/Stop test
	//ToDo: Add Destroy test
}
