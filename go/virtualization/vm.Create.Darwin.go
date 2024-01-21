package virtualization

import "github.com/sam-caldwell/monorepo/go/virtualization/vmDarwin"

// Create - Create a new VM blank VM in a stopped state (no opsys installed)
func (vm *VM) Create() *VM {

	configuration := vmDarwin.CreateVirtualMachineConfiguration(int(vm.cpuCount), int64(vm.ram))

	vmDarwin.CreateVirtualMachine(&configuration, vm.name.Get(), vm.name.Get())

	return vm
}
