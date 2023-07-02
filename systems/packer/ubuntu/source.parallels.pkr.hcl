#
# systems/packer/ubuntu/source.parallels.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
source "parallels-iso" "ubuntu" {
  iso_url = var.iso_url
  iso_checksum = var.iso_checksum
  iso_checksum_type = var.iso_checksum_type

  guest_os_type = "ubuntu"

  disk_size = var.disk_size

  ssh_username = var.username
  ssh_password = var.password
}