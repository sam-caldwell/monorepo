package virtualization

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/types/size"
	"time"
)

func NewVm() *VM {
	var vm VM

	// Set the default values
	_ = vm.name.Set(fmt.Sprintf("vmName%d", time.Now().UnixNano()))
	vm.ram = 4 * size.GB
	vm.cpuCount = 1

	// Return the address of the new VM object
	return &vm
}
