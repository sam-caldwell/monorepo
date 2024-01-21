package virtualization

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmNet"
)

// NetworkInterface - add a new network interface to the virtual machine definition
func (vm *VM) NetworkInterface(name string, interfaceType vmNet.InterfaceType) *VM {
	if vm.readOnly {

		vm.errors.Push(fmt.Errorf("readonly violation (NetworkInterface)"))

	} else {

		var net vmNet.NetworkInterfaceDescriptor

		if err := net.Name(name); err != nil {
			vm.errors.Push(err)
		}

		if err := net.Type(interfaceType); err != nil {
			vm.errors.Push(err)
		}

		vm.networkInterfaces = append(vm.networkInterfaces, net)

	}

	return vm

}
