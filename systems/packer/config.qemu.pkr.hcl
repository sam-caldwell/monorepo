locals {

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
}