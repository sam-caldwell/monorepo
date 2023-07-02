#
# systems/packer/ubuntu/source.qemu.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
source "qemu" "ubuntu" {
  iso_url = var.iso_url
  iso_checksum = var.iso_checksum
  iso_checksum_type = var.iso_checksum_type

  disk_size = "20000"

  http_directory = "http"

  ssh_username = var.username
  ssh_password = var.password
}