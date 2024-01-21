package virtualization

import "github.com/sam-caldwell/monorepo/go/types/size"

// Ram - Set the amount of RAM allocated to the virtual machine
func (vm *VM) Ram(sz size.Memory) *VM {
	vm.ram = sz
	return vm
}
