locals {
  communicator = var.communicator == null ? (
  var.is_windows ? "winrm" : "ssh"
  ) : var.communicator
}