package virtualization

// GetImage - return the virtual machine image file/path
func (vm *VM) GetImage() string {

	return vm.image.Get()

}
