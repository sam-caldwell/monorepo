package virtualization

// GetErrors - return a list of errors in the VM object
func (vm *VM) GetErrors() []error {
	return vm.errors.Dump()
}
