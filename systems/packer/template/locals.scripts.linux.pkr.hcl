locals {

  debian_scripts = [
    "${path.root}/scripts/${var.os_name}/update_${var.os_name}.sh",
    "${path.root}/scripts/_common/motd.sh",
    "${path.root}/scripts/_common/sshd.sh",
    "${path.root}/scripts/${var.os_name}/networking_${var.os_name}.sh",
    "${path.root}/scripts/${var.os_name}/sudoers_${var.os_name}.sh",
    "${path.root}/scripts/_common/vagrant.sh",
    "${path.root}/scripts/${var.os_name}/systemd_${var.os_name}.sh",
    "${path.root}/scripts/_common/virtualbox.sh",
    "${path.root}/scripts/_common/vmware_debian_ubuntu.sh",
    "${path.root}/scripts/_common/parallels.sh",
    "${path.root}/scripts/${var.os_name}/cleanup_${var.os_name}.sh",
    "${path.root}/scripts/_common/minimize.sh"
  ]

  default_linux_scripts = [
    "${path.root}/scripts/rhel/update_dnf.sh",
    "${path.root}/scripts/_common/motd.sh",
    "${path.root}/scripts/_common/sshd.sh",
    "${path.root}/scripts/_common/vagrant.sh",
    "${path.root}/scripts/_common/virtualbox.sh",
    "${path.root}/scripts/_common/vmware_rhel.sh",
    "${path.root}/scripts/_common/parallels-rhel.sh",
    "${path.root}/scripts/rhel/cleanup_dnf.sh",
    "${path.root}/scripts/_common/minimize.sh"
  ]

  fedora_scripts = [
    "${path.root}/scripts/fedora/networking_fedora.sh",
    "${path.root}/scripts/fedora/update_dnf.sh",
    "${path.root}/scripts/fedora/build-tools_fedora.sh",
    "${path.root}/scripts/fedora/install-supporting-packages_fedora.sh",
    "${path.root}/scripts/_common/motd.sh",
    "${path.root}/scripts/_common/sshd.sh",
    "${path.root}/scripts/_common/virtualbox.sh",
    "${path.root}/scripts/_common/vmware_fedora.sh",
    "${path.root}/scripts/_common/parallels-rhel.sh",
    "${path.root}/scripts/_common/vagrant.sh",
    "${path.root}/scripts/fedora/real-tmp_fedora.sh",
    "${path.root}/scripts/fedora/cleanup_dnf.sh",
    "${path.root}/scripts/_common/minimize.sh"
  ]

}

locals {
  is_linux_debian = var.os_name == "ubuntu" || var.os_name == "debian"
  is_linux_fedora = var.os_name == "fedora" || var.os_name == "redhat" || var.os_name == "centos"
}

locals {
  is_linux = local.is_linux_debian || local.is_linux_fedora
}

locals {
  linux_scripts = local.is_linux_debian ? local.debian_scripts : (
  local.is_linux_fedora ? local.fedora_scripts : local.default_linux_scripts
  )
}
