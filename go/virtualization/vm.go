package virtualization

import (
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/types/alphaNumericIdentifier"
	"github.com/sam-caldwell/monorepo/go/types/errorQueue"
	"github.com/sam-caldwell/monorepo/go/types/size"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmNet"
)

// VM - An object that represents a single virtual machine
type VM struct {
	name              alphaNumericIdentifier.Identifier // vm name.
	image             file.File                         // file and path to the resulting image.
	iso               file.File                         // file and path to the input ISO image.
	cpuCount          uint                              // number CPU cores.
	ram               size.Memory
	disks             []vmDisk.DiskDescriptor
	networkInterfaces []vmNet.NetworkInterfaceDescriptor
	errors            errorQueue.ErrorQueue
}
