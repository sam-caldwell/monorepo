package virtualization

// Cpu - set the number of CPU cores for the virtual machine
func (vm *VM) Cpu(n uint) *VM {
	vm.cpuCount = n
	return vm
}
