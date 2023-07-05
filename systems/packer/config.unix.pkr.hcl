locals {
  #  scripts = var.opsys == "solaris" ? [
  #    "${local.script_dir}/solaris/update_solaris.sh",
  #    "${local.script_dir}/_common/vagrant.sh",
  #    "${local.script_dir}/solaris/vmtools_solaris.sh",
  #    "${local.script_dir}/solaris/minimize_solaris.sh"
  #  ] : (
  #  var.opsys == "freebsd" ? [
  #    "${local.script_dir}/freebsd/update_freebsd.sh",
  #    "${local.script_dir}/freebsd/postinstall_freebsd.sh",
  #    "${local.script_dir}/freebsd/sudoers_freebsd.sh",
  #    "${local.script_dir}/_common/vagrant.sh",
  #    "${local.script_dir}/freebsd/vmtools_freebsd.sh",
  #    "${local.script_dir}/freebsd/cleanup_freebsd.sh",
  #    "${local.script_dir}/freebsd/minimize_freebsd.sh"
  #  ]
}