locals {
  scripts = local.fact.is_windows ? (
  substr(var.os_version, 0, 2) == "10" ||
  substr(var.os_version, 0, 2) == "11" ? [
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
  ] : [
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
  ) : (
  var.os_name == "solaris" ? [
    "${path.root}/scripts/solaris/update_solaris.sh",
    "${path.root}/scripts/_common/vagrant.sh",
    "${path.root}/scripts/solaris/vmtools_solaris.sh",
    "${path.root}/scripts/solaris/minimize_solaris.sh"
  ] : (
  var.os_name == "freebsd" ? [
    "${path.root}/scripts/freebsd/update_freebsd.sh",
    "${path.root}/scripts/freebsd/postinstall_freebsd.sh",
    "${path.root}/scripts/freebsd/sudoers_freebsd.sh",
    "${path.root}/scripts/_common/vagrant.sh",
    "${path.root}/scripts/freebsd/vmtools_freebsd.sh",
    "${path.root}/scripts/freebsd/cleanup_freebsd.sh",
    "${path.root}/scripts/freebsd/minimize_freebsd.sh"
  ] : (
  var.os_name == "opensuse" ||
  var.os_name == "sles" ? [
    "${path.root}/scripts/suse/repositories_suse.sh",
    "${path.root}/scripts/suse/update_suse.sh",
    "${path.root}/scripts/_common/motd.sh",
    "${path.root}/scripts/_common/sshd.sh",
    "${path.root}/scripts/_common/vagrant.sh",
    "${path.root}/scripts/suse/unsupported-modules_suse.sh",
    "${path.root}/scripts/_common/virtualbox.sh",
    "${path.root}/scripts/_common/vmware_suse.sh",
    "${path.root}/scripts/_common/parallels.sh",
    "${path.root}/scripts/suse/vagrant_group_suse.sh",
    "${path.root}/scripts/suse/sudoers_suse.sh",
    "${path.root}/scripts/suse/zypper-locks_suse.sh",
    "${path.root}/scripts/suse/remove-dvd-source_suse.sh",
    "${path.root}/scripts/suse/cleanup_suse.sh",
    "${path.root}/scripts/_common/minimize.sh"
  ] : (
  var.os_name == "ubuntu" ||
  var.os_name == "debian" ? [
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
    "${path.root}/scripts/${var.os_name}/hyperv_${var.os_name}.sh",
    "${path.root}/scripts/${var.os_name}/cleanup_${var.os_name}.sh",
    "${path.root}/scripts/_common/minimize.sh"
  ] : (
  var.os_name == "fedora" ? [
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
  ] : (
  "${var.os_name}-${substr(var.os_version, 0, 1)}" == "amazonlinux-2" ||
  "${var.os_name}-${substr(var.os_version, 0, 1)}" == "centos-7" ||
  "${var.os_name}-${substr(var.os_version, 0, 1)}" == "oracle-7" ||
  "${var.os_name}-${substr(var.os_version, 0, 1)}" == "rhel-7" ||
  "${var.os_name}-${substr(var.os_version, 0, 1)}" == "scientificlinux-7" ||
  "${var.os_name}-${substr(var.os_version, 0, 1)}" == "springdalelinux-7" ? [
    "${path.root}/scripts/rhel/update_yum.sh",
    "${path.root}/scripts/_common/motd.sh",
    "${path.root}/scripts/_common/sshd.sh",
    "${path.root}/scripts/rhel/networking_rhel7.sh",
    "${path.root}/scripts/_common/vagrant.sh",
    "${path.root}/scripts/_common/virtualbox.sh",
    "${path.root}/scripts/_common/vmware_rhel.sh",
    "${path.root}/scripts/_common/parallels-rhel.sh",
    "${path.root}/scripts/rhel/cleanup_yum.sh",
    "${path.root}/scripts/_common/minimize.sh"
  ] : [
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
  )
  )
  )
  )
  )
  )
}