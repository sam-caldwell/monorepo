#
# systems/packer/source.vmware.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
#source "vmware-iso" "vm" {
#  vm_name                = var.prefix + var.opsys
#
#  shutdown_command  = locals[os_family][os_version][shutdown]
#
#  iso_urls          = locals[os_family][os_version][iso]
#  iso_checksum      = locals[os_family][os_version][iso_checksum]
#  iso_checksum_type = locals.fact.iso_checksum_type
#
#  guest_os_type = locals[os_family][os_version][iso_checksum]
#
#  disk_size = locals.fact.size
#
#  communicator   = locals[os_family][os_version][communicator]
#  winrm_username = locals.fact.user.username
#  winrm_password = locals.fact.user.password
#}
