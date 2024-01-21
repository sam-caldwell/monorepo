package virtualization

import "fmt"

// Cpu - set the number of CPU cores for the virtual machine
func (vm *VM) Cpu(n uint) *VM {
	if vm.readOnly {
		vm.errors.Push(fmt.Errorf("readonly violation (cpu)"))
	} else {
		vm.cpuCount = n
	}
	return vm
}
