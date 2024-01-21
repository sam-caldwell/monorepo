package virtualization

// Iso - Specify the filename/path of the iso image.
func (vm *VM) Iso(fileName string) *VM {
	if err := vm.iso.SetWithExtension(fileName, []string{".iso"}); err != nil {
		vm.errors.Push(err)
	}
	return vm
}
