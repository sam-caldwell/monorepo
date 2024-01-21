package vmNet

import "github.com/sam-caldwell/monorepo/go/types/alphaNumericIdentifier"

type NetworkInterfaceDescriptor struct {
	name          alphaNumericIdentifier.Identifier
	interfaceType InterfaceType
}
