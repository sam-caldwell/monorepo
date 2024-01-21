package virtualization

import "fmt"

// Name - Set the virtual machine name and return *VM to allow method chaining.
func (vm *VM) Name(name string) *VM {

	if vm.readOnly {
		vm.errors.Push(fmt.Errorf("readonly violation (name)"))
	} else {
		if err := vm.name.Set(name); err != nil {

			vm.errors.Push(err)

			//Note: Name will be set to a default pattern by the Set() method

		}
	}
	return vm

}
