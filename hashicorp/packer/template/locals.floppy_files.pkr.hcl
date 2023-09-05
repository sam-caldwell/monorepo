locals {
  default_windows_floppy_files = [
    "${path.root}/win_answer_files/${var.os_version}/Autounattend.xml",
    "${path.root}/scripts/windows/base_setup.ps1"
  ]
  default_linux_floppy_files = []
}

locals {
  floppy_files = var.floppy_files == null ? (
  var.is_windows ? local.default_windows_floppy_files : local.default_linux_floppy_files
  ) : var.floppy_files
}
