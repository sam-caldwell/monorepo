package virtualization

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/types/size"
)

// Ram - Set the amount of RAM allocated to the virtual machine
func (vm *VM) Ram(sz size.Memory) *VM {
	if vm.readOnly {
		vm.errors.Push(fmt.Errorf("readonly violation (ram)"))
	} else {
		vm.ram = sz
	}
	return vm
}
