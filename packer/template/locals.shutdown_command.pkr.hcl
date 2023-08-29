locals {
  default_windows_shutdown = "shutdown /s /t 10 /f /d p:4:1 /c \"Packer Shutdown\""
  default_linux_shutdown   = "echo 'vagrant' | sudo -S /sbin/halt -h -p"
}

locals {
  shutdown_command = var.shutdown_command == null ? (
  var.is_windows ? local.default_windows_shutdown : local.default_linux_shutdown
  ) : var.shutdown_command
}