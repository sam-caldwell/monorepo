source "virtualbox-iso" "vm" {
  # Virtualbox specific options
  gfx_controller            = local.vbox_gfx_controller
  gfx_vram_size             = local.vbox_gfx_vram_size
  guest_additions_path      = var.vbox_guest_additions_path
  guest_additions_mode      = local.vbox_guest_additions_mode
  guest_additions_interface = var.vbox_guest_additions_interface
  guest_os_type             = var.vbox_guest_os_type
  hard_drive_interface      = var.vbox_hard_drive_interface
  iso_interface             = var.vbox_iso_interface
  vboxmanage                = var.vboxmanage
  virtualbox_version_file   = var.virtualbox_version_file
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
  output_directory = "${local.output_directory}-virtualbox"
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
source "virtualbox-ovf" "amazonlinux" {
  # Virtualbox specific options
  guest_additions_path    = var.vbox_guest_additions_path
  source_path             = local.vbox_source
  vboxmanage              = var.vboxmanage
  virtualbox_version_file = var.virtualbox_version_file
  # Source block common options
  communicator     = local.communicator
  headless         = var.headless
  output_directory = "${local.output_directory}-virtualbox-ovf"
  shutdown_command = local.shutdown_command
  shutdown_timeout = var.shutdown_timeout
  ssh_password     = var.ssh_password
  ssh_port         = var.ssh_port
  ssh_timeout      = var.ssh_timeout
  ssh_username     = var.ssh_username
  vm_name          = local.vm_name
}