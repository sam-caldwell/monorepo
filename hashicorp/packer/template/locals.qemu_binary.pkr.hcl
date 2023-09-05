locals {

  qemu_binary = var.qemu_binary == null ? "qemu-system-${var.os_arch}" : var.qemu_binary

}