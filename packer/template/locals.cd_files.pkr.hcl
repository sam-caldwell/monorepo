locals {
  default_windows_cd_files = "${path.root}/win_answer_files/${substr(var.os_version, 0, 2)}/gen2_Autounattend.xml"
  default_linux_cd_files   = null
}

locals {
  cd_files = var.cd_files == null ? (
  var.is_windows ? [
    local.default_windows_cd_files
  ] : local.default_linux_cd_files
  ) : var.cd_files
}