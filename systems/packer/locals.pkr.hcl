/*
 * This file contains the calculated locals using variables and other
 * locals files.
 */
locals {
  // Resolve our hyper-v configuration...
  hyperv = {

    generation = 2

  }


  // Resolve our qemu configuration...
  #  qemu = {
  #
  #    qemu_binary = var.qemu_binary == null ? "qemu-system-${var.os_arch}" : var.qemu_binary
  #
  #    qemu_machine_type = var.qemu_machine_type == null ? (
  #    var.os_arch == "aarch64" ? "virt" : "q35"
  #    ) : var.qemu_machine_type
  #
  #    qemuargs = local.hyperv.generation == 2 && local[var.os_family][var.opsys][var.os_version].is_windows ? [
  #      ["-bios", "/usr/share/OVMF/OVMF_CODE.fd"],
  #    ] : (
  #    local[var.os_family][var.opsys][var.os_version].is_windows ? [
  #      ["-drive", "file=${path.root}/win_answer_files/virtio-win.iso,media=cdrom,index=3"],
  #      [
  #        "-drive",
  #        "file=${path.root}/../builds/packer-${var.os_name}-${var.os_version}-${var.os_arch}-qemu/{{ .Name }},if=virtio,cache=writeback,discard=ignore,format=qcow2,index=1"
  #      ],
  #    ] : (
  #    local[var.os_family][var.opsys][var.os_version].arch == "aarch64" ? [
  #      ["-boot", "strict=off"]
  #    ] : null
  #    )
  #    )
  #    ) : var.qemuargs
  #  }
  // Resolve our virtual box configuration...
  #  vbox= {
  #    /*
  #   * VIRTUALBOX: vbox_gfx_controller
  #   */
  #    vbox_gfx_controller = var.vbox_gfx_controller == null ? (
  #    local.fact.is_windows ? "vboxsvga" : "vmsvga"
  #    ) : var.vbox_gfx_controller
  #    /*
  #   * VIRTUALBOX: vbox_gfx_vram_size
  #   */
  #    vbox_gfx_vram_size = var.vbox_gfx_controller == null ? (
  #    local.fact.is_windows ? 128 : 33
  #    ) : var.vbox_gfx_vram_size
  #    /*
  #   * VIRTUALBOX: vbox_guest_additions_mode
  #   */
  #    vbox_guest_additions_mode = var.vbox_guest_additions_mode == null ? (
  #    local.fact.is_windows && var.hyperv_generation == 1 ? "attach" : "upload"
  #    ) : var.vbox_guest_additions_mode
  #    /*
  #   * VIRTUALBOX: vbox_source
  #   */
  #    # virtualbox-ovf
  #    vbox_source = var.vbox_source == null ? (
  #    var.os_name == "amazonlinux" ? "${path.root}/amz_working_files/amazon2.ovf" : null
  #    ) : var.vbox_source
  #  }
  // Resolve our vmware configuration...
  vmware = {

  }

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
#  http_directory   = var.http_directory == null ? "${path.root}/http" : var.http_directory
#  memory           = var.memory == null ? (local.fact.is_windows ? 4096 : 2048) : var.memory
#  shutdown_command = var.shutdown_command == null ? (
#  local.fact.is_windows ? "shutdown /s /t 10 /f /d p:4:1 /c \"Packer Shutdown\"" : (
#  var.os_name == "freebsd" ? "echo 'vagrant' | su -m root -c 'shutdown -p now'" : "echo 'vagrant' | sudo -S /sbin/halt -h -p"
#  )
#  ) : var.shutdown_command
}
