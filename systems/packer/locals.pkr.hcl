/*
 * This file contains the calculated locals using variables and other
 * locals files.
 */
locals {

  // Resolve our vmware configuration...


  # vmware-iso

  # Source block common

  #  cd_files = var.cd_files == null ? (
  #  var.hyperv_generation == 2 && local.fact.is_windows ? [
  #    "${path.root}/win_answer_files/${substr(var.os_version, 0, 2)}/gen2_Autounattend.xml"
  #  ] : null
  #  ) : var.cd_files
  #  communicator = var.communicator == null ? (
  #  local.fact.is_windows ? "winrm" : "ssh"
  #  ) : var.communicator
  #  floppy_files = var.floppy_files == null ? (
  #  var.hyperv_generation == 2 ? null : (
  #  local.fact.is_windows ? [
  #    "${path.root}/win_answer_files/${var.os_version}/Autounattend.xml",
  #    "${path.root}/scripts/windows/base_setup.ps1"
  #  ] : (
  #  var.os_name == "springdalelinux" ? [
  #    "${path.root}/http/rhel/${substr(var.os_version, 0, 1)}ks.cfg"
  #  ] : null
  #  )
  #  )
  #  ) : var.floppy_files
  #  shutdown_command = var.shutdown_command == null ? (
  #  local.fact.is_windows ? "shutdown /s /t 10 /f /d p:4:1 /c \"Packer Shutdown\"" : (
  #  var.os_name == "freebsd" ? "echo 'vagrant' | su -m root -c 'shutdown -p now'" : "echo 'vagrant' | sudo -S /sbin/halt -h -p"
  #  )
  #  ) : var.shutdown_command
}
