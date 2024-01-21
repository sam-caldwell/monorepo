package vmDisk

import (
	"github.com/sam-caldwell/monorepo/go/types/alphaNumericIdentifier"
	"github.com/sam-caldwell/monorepo/go/types/size"
	"github.com/sam-caldwell/monorepo/go/virtualization/vmDisk/diskFormat"
)

type DiskDescriptor struct {
	name   alphaNumericIdentifier.Identifier
	size   size.Memory
	format diskFormat.DiskFormat
	//ToDo: Add other things...?
}
