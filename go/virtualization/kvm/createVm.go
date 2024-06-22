package kvm

import (
	"libvirt.org/go/libvirt"
)

// CreateVm - Launch a defined domain.
//
// See https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainCreate
//
// Launch a defined domain. If the call succeeds the domain moves from the defined to the running domains pools.
// The domain will be paused only if restoring from managed state created from a paused domain. For more control,
// see virDomainCreateWithFlags().
//
//	(c) 2023 Sam Caldwell.  MIT License
func CreateVm(domain *libvirt.Domain) error {
	return domain.Create()
}
