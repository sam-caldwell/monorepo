locals {
  boot_command = [
    join("", var.boot_command)
  ]
}