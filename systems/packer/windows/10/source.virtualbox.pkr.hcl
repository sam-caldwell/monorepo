#
# systems/packer/windows10/source.virtualbox.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
source "virtualbox-iso" "windows10" {
  iso_url = var.iso_url
  iso_checksum = var.iso_checksum
  iso_checksum_type = var.iso_checksum_type
  guest_os_type = "Windows10_64"
  headless = true

  disk_size = var.disk_size

  communicator = "winrm"
  winrm_username = var.username
  winrm_password = var.password
}