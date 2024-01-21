package virtualization

import "fmt"

// Image - Set the virtual machine image filename/path and return *VM to allow method chaining.
func (vm *VM) Image(name string) *VM {

	if vm.readOnly {
		vm.errors.Push(fmt.Errorf("readonly violation (image)"))
	} else {
		if err := vm.image.Set(name); err != nil {

			vm.errors.Push(err)

		}
	}
	return vm

}
