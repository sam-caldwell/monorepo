source "parallels-iso" "vm" {
  boot_command           = var.boot_command
  boot_wait              = local.boot_wait
  cpus                   = var.cpus
  communicator           = local.communicator
  disk_size              = var.disk_size
  floppy_files           = local.floppy_files
  guest_os_type          = var.parallels_guest_os_type
  http_directory         = local.http_directory
  iso_checksum           = var.iso_checksum
  iso_url                = var.iso_url
  memory                 = local.memory
  output_directory       = "${local.output_directory}-parallels"
  parallels_tools_flavor = local.parallels_tools_flavor
  parallels_tools_mode   = local.parallels_tools_mode
  prlctl                 = local.parallels_prlctl
  prlctl_version_file    = var.parallels_prlctl_version_file
  shutdown_command       = local.shutdown_command
  shutdown_timeout       = var.shutdown_timeout
  ssh_password           = var.ssh_password
  ssh_port               = var.ssh_port
  ssh_timeout            = var.ssh_timeout
  ssh_username           = var.ssh_username
  winrm_password         = var.winrm_password
  winrm_timeout          = var.winrm_timeout
  winrm_username         = var.winrm_username
  vm_name                = local.vm_name
}