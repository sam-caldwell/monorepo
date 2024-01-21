package virtualization

// GetName - Return the virtual machine name (not a chainable method)
func (vm *VM) GetName() string {

	return vm.name.Get()

}
