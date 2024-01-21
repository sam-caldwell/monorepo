package virtualization

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/environment"
	"strings"
)

// Image - Set the virtual machine image filename/path and return *VM to allow method chaining.
func (vm *VM) Image(name string) *VM {

	if vm.readOnly {

		vm.errors.Push(fmt.Errorf("readonly violation (image)"))

	} else {

		envHome, err := environment.RequireString("HOME")

		if err != nil {
			vm.errors.Push(err)
		}

		name = strings.Replace(name, "{{name}}", vm.GetName(), -1)

		name = strings.Replace(name, "{{home}}", envHome, -1)

		if err := vm.image.Set(name); err != nil {

			vm.errors.Push(err)

		}
	}

	return vm

}
