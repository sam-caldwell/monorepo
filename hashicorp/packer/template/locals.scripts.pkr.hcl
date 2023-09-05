locals {
  scripts = var.is_windows ? local.windows_scripts : (local.is_linux ? local.linux_scripts : [])
}