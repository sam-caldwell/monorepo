locals{
  cpu_arch = var.os_arch == "x86_64"?"amd64":var.os_arch
}

locals {
  vm_name = var.vm_name == null ? (
  var.os_arch == "${var.os_name}-${var.os_version}-${local.cpu_arch}"
  ) : var.vm_name
}