//go:build windows
// +build windows

package virtualization

func (vm *VM) Create() {
	vm.errors.Push(fmt.Errorf("windows is not implemented"))
}
