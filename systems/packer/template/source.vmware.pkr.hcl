source "vmware-iso" "vm" {
  # VMware specific options
  cdrom_adapter_type             = var.vmware_cdrom_adapter_type
  disk_adapter_type              = var.vmware_disk_adapter_type
  guest_os_type                  = var.vmware_guest_os_type
  network                        = var.vmware_network
  network_adapter_type           = var.vmware_network_adapter_type
  tools_upload_flavor            = local.vmware_tools_upload_flavor
  tools_upload_path              = local.vmware_tools_upload_path
  usb                            = var.vmware_enable_usb
  version                        = var.vmware_version
  vmx_data                       = var.vmware_vmx_data
  vmx_remove_ethernet_interfaces = var.vmware_vmx_remove_ethernet_interfaces
  # Source block common options
  boot_command     = var.boot_command
  boot_wait        = local.boot_wait
  cpus             = var.cpus
  communicator     = local.communicator
  disk_size        = var.disk_size
  floppy_files     = local.floppy_files
  headless         = var.headless
  http_directory   = local.http_directory
  iso_checksum     = var.iso_checksum
  iso_url          = var.iso_url
  memory           = local.memory
  output_directory = "${local.output_directory}-vmware"
  shutdown_command = local.shutdown_command
  shutdown_timeout = var.shutdown_timeout
  ssh_password     = var.ssh_password
  ssh_port         = var.ssh_port
  ssh_timeout      = var.ssh_timeout
  ssh_username     = var.ssh_username
  winrm_password   = var.winrm_password
  winrm_timeout    = var.winrm_timeout
  winrm_username   = var.winrm_username
  vm_name          = local.vm_name
}