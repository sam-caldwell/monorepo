#
# systems/packer/ubuntu/source.vmware.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
source "vmware-iso" "ubuntu" {
  iso_url = var.iso_url
  iso_checksum = var.iso_checksum
  iso_checksum_type = var.iso_checksum_type

  guest_os_type = "ubuntu-64"

  disk_size = var.disk_size

  ssh_username = var.username
  ssh_password = var.password
}
