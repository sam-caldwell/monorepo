locals {
  /*
   * compression level = {0-9}
   * Higher compression means it takes longer to complete.
   */
  vagrant_compression_level = 9
  /*
   * The name of the output vagrant box.
   */
  vagrant_output = join("", [
    "build/vagrant/boxes/", "{{.Provider}}", ".", var.opsys, ".", var.os_version, ".box"
  ])
  /*
   *
   */
  vagrant_template = local.fact.is_windows ? (
    local.hyperv.generation == 1 ?
    "${path.root}/scripts/vagrantfile-windows.template" :
    "${path.root}/scripts/vagrantfile-windows-gen2.template"
  ) : null
}