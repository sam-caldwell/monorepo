package vmNet

import "github.com/sam-caldwell/monorepo/go/types/alphaNumericIdentifier"

type NetworkInterfaceDescriptor struct {
	name          alphaNumericIdentifier.Identifier
	interfaceType NetworkInterfaceType
}

func (iface *NetworkInterfaceDescriptor) GetName() string {
	return iface.name.Get()
}

func (iface *NetworkInterfaceDescriptor) GetType() NetworkInterfaceType {
	return iface.interfaceType
}

func (iface *NetworkInterfaceDescriptor) Name(n string) error {
	return iface.name.Set(n)
}

func (iface *NetworkInterfaceDescriptor) Size(interfaceType NetworkInterfaceType) error {
	iface.interfaceType = interfaceType
	return nil
}
