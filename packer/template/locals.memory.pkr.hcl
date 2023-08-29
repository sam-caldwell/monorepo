locals {
  default_windows_memory = 4096
  default_linux_memory   = 2048
}
locals {
  memory = var.memory == null ? (
  var.is_windows ? local.default_windows_memory : local.default_linux_memory
  ) : var.memory
}