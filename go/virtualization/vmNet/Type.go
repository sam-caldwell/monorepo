package vmNet

func (iface *NetworkInterfaceDescriptor) Type(interfaceType InterfaceType) error {
	iface.interfaceType = interfaceType
	return nil
}
