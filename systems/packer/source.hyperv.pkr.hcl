/*
 * systems/packer/source.hyperv.pkr.hcl
 * (c) Sam Caldwell.  See LICENSE.txt
 */

#source "hyperv-iso" "vm" {
  # Hyper-v specific options
#  enable_dynamic_memory = local.hyperv_enable_dynamic_memory
#  enable_secure_boot    = local.hyperv_enable_secure_boot
#  generation            = var.hyperv_generation
#  guest_additions_mode  = var.hyperv_guest_additions_mode
#  switch_name           = var.hyperv_switch_name
#  # Source block common options
#  boot_command          = var.boot_command
#  boot_wait             = local.boot_wait
#  cpus                  = var.cpus
#  communicator          = local.communicator
#  disk_size             = var.disk_size
#  floppy_files          = local.floppy_files
#  headless              = var.headless
#  http_directory        = local.http_directory
#  iso_checksum          = var.iso_checksum
#  iso_url               = var.iso_url
#  memory                = local.memory
#  output_directory      = "${local.output_directory}-hyperv"
#  shutdown_command      = local.shutdown_command
#  shutdown_timeout      = var.shutdown_timeout
#  ssh_password          = var.ssh_password
#  ssh_port              = var.ssh_port
#  ssh_timeout           = var.ssh_timeout
#  ssh_username          = var.ssh_username
#  winrm_password        = var.winrm_password
#  winrm_timeout         = var.winrm_timeout
#  winrm_username        = var.winrm_username
#  vm_name               = local.vm_name
#}