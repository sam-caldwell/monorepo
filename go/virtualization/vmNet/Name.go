package vmNet

func (iface *NetworkInterfaceDescriptor) Name(n string) error {
	return iface.name.Set(n)
}
