locals {

  windows_10_scripts = [
    # "${path.root}/scripts/windows/base_setup.ps1",
    "${path.root}/scripts/windows/provision.ps1",
    "${path.root}/scripts/windows/disable-windows-updates.ps1",
    "${path.root}/scripts/windows/disable-windows-defender.ps1",
    "${path.root}/scripts/windows/remove-one-drive.ps1",
    "${path.root}/scripts/windows/remove-apps.ps1",
    "${path.root}/scripts/windows/virtualbox-prevent-vboxsrv-resolution-delay.ps1",
    "${path.root}/scripts/windows/provision-winrm.ps1",
    "${path.root}/scripts/windows/enable-remote-desktop.ps1",
    "${path.root}/scripts/windows/eject-media.ps1"
  ]
  default_windows_scripts = [
    # "${path.root}/scripts/windows/base_setup.ps1",
    "${path.root}/scripts/windows/provision.ps1",
    "${path.root}/scripts/windows/disable-windows-updates.ps1",
    "${path.root}/scripts/windows/disable-windows-defender.ps1",
    "${path.root}/scripts/windows/remove-one-drive.ps1",
    # "${path.root}/scripts/windows/remove-apps.ps1",
    "${path.root}/scripts/windows/virtualbox-prevent-vboxsrv-resolution-delay.ps1",
    "${path.root}/scripts/windows/provision-winrm.ps1",
    "${path.root}/scripts/windows/enable-remote-desktop.ps1",
    "${path.root}/scripts/windows/eject-media.ps1"
  ]

}

locals {

  windows_os_version = substr(var.os_version, 0, 2)
  is_windows10       = var.is_windows && (local.windows_os_version == "10")
  is_windows11       = var.is_windows && (local.windows_os_version == "11")
}

locals {
  windows_scripts = var.is_windows ?((local.is_windows10 || local.is_windows11)
  ? local.windows_10_scripts
  : local.default_windows_scripts) : null
}