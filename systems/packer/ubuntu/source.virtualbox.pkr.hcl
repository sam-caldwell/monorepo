#
# systems/packer/ubuntu/source.virtualbox.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
source "virtualbox-iso" "ubuntu" {
  iso_url = var.iso_url
  iso_checksum = var.iso_checksum
  iso_checksum_type = var.iso_checksum_type

  guest_os_type = "Ubuntu_64"

  disk_size = var.disk_size

  ssh_username = var.username
  ssh_password = var.password
  headless = true
}
