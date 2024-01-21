package vmNet

import "github.com/sam-caldwell/monorepo/go/types/alphaNumericIdentifier"

type NetworkInterfaceDescriptor struct {
	name          alphaNumericIdentifier.Identifier
	interfaceType InterfaceType
}

func (iface *NetworkInterfaceDescriptor) GetName() string {
	return iface.name.Get()
}

func (iface *NetworkInterfaceDescriptor) GetType() InterfaceType {
	return iface.interfaceType
}

func (iface *NetworkInterfaceDescriptor) Name(n string) error {
	return iface.name.Set(n)
}

func (iface *NetworkInterfaceDescriptor) Type(interfaceType InterfaceType) error {
	iface.interfaceType = interfaceType
	return nil
}
