#
# systems/packer/windows10/source.qemu.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
source "qemu" "windows10" {
  iso_url = var.iso_url
  iso_checksum = var.iso_checksum
  iso_checksum_type = var.iso_checksum_type

  disk_size = var.disk_size

  http_directory = "http"

  communicator = "winrm"
  winrm_username = var.username
  winrm_password = var.password
}
