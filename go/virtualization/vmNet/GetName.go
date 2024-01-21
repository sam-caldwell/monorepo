package vmNet

func (iface *NetworkInterfaceDescriptor) GetName() string {
	return iface.name.Get()
}
