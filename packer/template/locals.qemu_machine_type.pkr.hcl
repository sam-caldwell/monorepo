locals {

  qemu_machine_type = var.qemu_machine_type == null ? (
  var.os_arch == "aarch64" ? "virt" : "q35"
  ) : var.qemu_machine_type

}