#
# systems/packer/source.parallels.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#

source "parallels-iso" "vm" {
  boot_wait              = local[var.os_family][var.opsys][var.os_version].boot_wait
  boot_command           = local[var.os_family][var.opsys][var.os_version].boot_command
  cpus                   = local.fact.cpus
  communicator           = local[var.os_family][var.opsys][var.os_version].communicator
  disk_size              = local.fact.disk.size
  floppy_files           = local[var.os_family][var.opsys][var.os_version].floppy_files
  guest_os_type          = local[var.os_family][var.opsys][var.os_version].guest_os_type.parallels
  http_directory         = local.fact.http_directory
  iso_checksum           = local[var.os_family][var.opsys][var.os_version].iso_checksum
  #  iso_checksum_type = local.fact.iso_checksum_type
  iso_urls               = [local[var.os_family][var.opsys][var.os_version].iso]
  memory                 = local.fact.memory
  output_directory       = local.output_directory
  parallels_tools_flavor = local.parallels.tool.flavors[local[var.os_family][var.opsys][var.os_version].arch][var.os_family]
  parallels_tools_mode   = local.fact.is_windows ? "attach" : "upload"
  prlctl                 = local.parallels.prlctl[var.os_family]
  prlctl_version_file    = local.parallels.prlctl.version_file
  shutdown_command       = local[var.os_family][var.opsys][var.os_version].shutdown.command
  shutdown_timeout       = local[var.os_family][var.opsys][var.os_version].shutdown.timeout
  ssh_password           = local.fact.user.password
  ssh_port               = local.fact.ssh.port
  ssh_timeout            = local.fact.ssh.timeout
  ssh_username           = local.fact.user.username
  vm_name                = local.vm_name
  winrm_password         = local.fact.user.password
  winrm_timeout          = local.fact.winrm.timeout
  winrm_username         = local.fact.user.username


  # Source block common options

}