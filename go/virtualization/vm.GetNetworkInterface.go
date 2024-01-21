package virtualization

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmNet"
)

// GetNetworkInterface - Return the given network interface descriptor, identified by the index
func (vm *VM) GetNetworkInterface(index uint) vmNet.NetworkInterfaceDescriptor {

	if index >= uint(len(vm.networkInterfaces)) {

		vm.errors.Push(fmt.Errorf("index out of range when calling VM.GetNetworkInterfaces()"))

		return vmNet.NetworkInterfaceDescriptor{}
	}

	return vm.networkInterfaces[index]

}
