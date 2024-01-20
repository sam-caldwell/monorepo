package vmname

// Get - Return the virtual machine identifier.
func (name *VmName) Get() string {
	return string(*name)
}
