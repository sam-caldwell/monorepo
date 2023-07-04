locals {
  #  scripts = var.os_name == "solaris" ? [
  #    "${path.root}/scripts/solaris/update_solaris.sh",
  #    "${path.root}/scripts/_common/vagrant.sh",
  #    "${path.root}/scripts/solaris/vmtools_solaris.sh",
  #    "${path.root}/scripts/solaris/minimize_solaris.sh"
  #  ] : (
  #  var.os_name == "freebsd" ? [
  #    "${path.root}/scripts/freebsd/update_freebsd.sh",
  #    "${path.root}/scripts/freebsd/postinstall_freebsd.sh",
  #    "${path.root}/scripts/freebsd/sudoers_freebsd.sh",
  #    "${path.root}/scripts/_common/vagrant.sh",
  #    "${path.root}/scripts/freebsd/vmtools_freebsd.sh",
  #    "${path.root}/scripts/freebsd/cleanup_freebsd.sh",
  #    "${path.root}/scripts/freebsd/minimize_freebsd.sh"
  #  ]
}