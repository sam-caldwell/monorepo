package virtualization

// GetIso - Return the current iso file/path.
func (vm *VM) GetIso() string {
	return vm.iso.Get()
}
