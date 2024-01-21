package vmNet

func (iface *NetworkInterfaceDescriptor) GetType() InterfaceType {
	return iface.interfaceType
}
