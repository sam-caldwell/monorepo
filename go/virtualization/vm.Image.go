package virtualization

// Image - Set the virtual machine image filename/path and return *VM to allow method chaining.
func (vm *VM) Image(name string) *VM {

	if err := vm.image.Set(name); err != nil {

		vm.errors.Push(err)

	}

	return vm

}
