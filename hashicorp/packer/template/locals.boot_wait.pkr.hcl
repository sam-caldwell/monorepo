locals {
  default_windows_boot_wait = "60s"
  default_linux_boot_wait   = "10s"
}

locals {
  boot_wait = var.boot_wait == null ? (
  var.is_windows ? local.default_windows_boot_wait : local.default_linux_boot_wait
  ) : var.boot_wait
}