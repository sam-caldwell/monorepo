#
# systems/packer/source.parallels.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#

source "parallels-iso" "vm" {

  vm_name = join("",[var.prefix, var.opsys, var.os_version])

  parallels_tools_flavor = "win" #local[var.os_family][var.os_version][var.guest_os_type].parallels

  boot_command = local[var.os_family][var.os_version].boot_command

  shutdown_command = local[var.os_family][var.os_version].shutdown

  iso_urls     = [local[var.os_family][var.os_version].iso]
  iso_checksum = local[var.os_family][var.os_version].iso_checksum
  #  iso_checksum_type = local.iso_checksum_type

  guest_os_type = local[var.os_family][var.os_version].iso_checksum

  disk_size = local.general.disk.size

  communicator   = local[var.os_family][var.os_version].communicator
  winrm_username = local.general.user.username
  winrm_password = local.general.user.password

}