
source "vmware-iso" "windows10" {
  iso_url = var.iso_url
  iso_checksum = var.iso_checksum
  iso_checksum_type = var.iso_checksum_type
  guest_os_type = "windows8srv-64"

  disk_size = var.disk_size

  communicator = "winrm"
  winrm_username = "vagrant"
  winrm_password = "vagrant"
}
