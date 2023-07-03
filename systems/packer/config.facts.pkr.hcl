/*
 * systems/packer/config.fact.pkr.hcl
 * (c) Sam Caldwell.  See LICENSE.txt
 *
 * This is a generic configuration file that
 * uses local variables to configure how our
 * packer builds will work.
 */
locals {
  fact = {
    cd_files = local.hyperv.generation == 2 && (var.os_family == "windows") ? [
      "${path.root}/win_answer_files/${var.os_version}/gen2_Autounattend.xml"
    ] : null
    floppy_files = local.hyperv.generation == 1 && (var.os_family == "windows") ? [
        "${path.root}/win_answer_files/${var.opsys}/${var.os_version}/Autounattend.xml",
        "${path.root}/scripts/${var.opsys}/base_setup.ps1"
      ] : null
    cpus = 4
    disk = {
      size = 20000 // disk space in MB
    }
    http_directory    = "${path.root}/http"
    http_proxy        = ""
    https_proxy       = ""
    hostname          = var.opsys
    no_proxy          = false
    is_windows        = var.os_family == "windows"
    iso_checksum_type = "sha256"
    memory            = 4096 // memory in MB
    output_directory  = "build/packer/{{.Provider}}/"
    source = {
      enabled = {
        hyperv     = (var.source == "hyperv") ? "sources.hyperv-iso.vm" : ""
        parallels  = (var.source == "parallels") ? "sources.parallels-iso.vm" : ""
        qemu       = (var.source == "qemu") ? "sources.qemu.vm" : ""
        virtualbox = (var.source == "virtualbox") ? "sources.virtualbox-iso.vm" : ""
        vmware     = (var.source == "vmware") ? "sources.vmware-iso.vm" : ""
      }
    }
    ssh = {
      port    = 2222 // Can we use port 0?
      timeout = "60s"
    }
    winrm = {
      timeout = "60s"
    }
    user = {
      username = "vagrant"
      password = "vagrant"
    }
  }
}
