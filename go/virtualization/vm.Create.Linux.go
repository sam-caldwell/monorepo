//go:build linux
// +build linux

package virtualization

func (vm *VM) Create() {
	vm.errors.Push(fmt.Errorf("linux is not implemented"))
}
