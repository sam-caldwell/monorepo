package virtualization

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/environment"
	"github.com/sam-caldwell/monorepo/go/fs/file/extensions"
	"strings"
)

// Iso - Specify the filename/path of the iso image.
func (vm *VM) Iso(fileName string) *VM {

	if vm.readOnly {

		vm.errors.Push(fmt.Errorf("readonly violation (iso)"))

	} else {

		envHome, err := environment.RequireString("HOME")

		if err != nil {
			vm.errors.Push(err)
		}

		fileName = strings.Replace(fileName, "{{name}}", vm.GetName(), -1)

		fileName = strings.Replace(fileName, "{{home}}", envHome, -1)

		if err := vm.iso.SetIfHasExtension(fileName, []string{extensions.ISO}); err != nil {
			vm.errors.Push(err)
		}
	}

	return vm

}
