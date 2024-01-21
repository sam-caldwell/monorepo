//go:build !darwin && !windows && !linux
// +build !darwin,!windows,!linux

package virtualization

import "fmt"

func (vm *VM) Create() {
	vm.errors.Push(fmt.Errorf("unsupported operating system"))
}
