/*
 * systems/packer/config.windows.pkr.hcl
 * (c) Sam Caldwell.  See LICENSE.txt
 *
 * Windows-specific configuration for windows 10.
 */
locals {
  //os_family
  windows = {
    //opsys
    windows = {
      //os_version
      10 = {
        arch          = "x86_64"
        boot_command  = []
        boot_wait     = "3600s"
        cd_files      = local.fact.cd_files
        communicator  = "winrm"
        floppy_files  = local.fact.floppy_files
        guest_os_type = {
          parallels = "win-10"
          vbox      = "Windows10_64"
          vmware    = "windows9srv-64"
        }
        iso          = "https://software-static.download.prss.microsoft.com/dbazure/988969d5-f34g-4e03-ac9d-1f9786c66750/19045.2006.220908-0225.22h2_release_svc_refresh_CLIENTENTERPRISEEVAL_OEMRET_x64FRE_en-us.iso"
        iso_checksum = "ef7312733a9f5d7d51cfa04ac497671995674ca5e1058d5164d6028f0938d668"
        scripts      = [
          # "${local.script_dir}/windows/base_setup.ps1",
          "${local.script_dir}/${var.os_family}/provision.ps1",
          "${local.script_dir}/${var.os_family}/disable-windows-updates.ps1",
          "${local.script_dir}/${var.os_family}/disable-windows-defender.ps1",
          "${local.script_dir}/${var.os_family}/remove-one-drive.ps1",
          "${local.script_dir}/${var.os_family}/remove-apps.ps1",
          "${local.script_dir}/${var.os_family}/virtualbox-prevent-vboxsrv-resolution-delay.ps1",
          "${local.script_dir}/${var.os_family}/provision-winrm.ps1",
          "${local.script_dir}/${var.os_family}/enable-remote-desktop.ps1",
          "${local.script_dir}/${var.os_family}/eject-media.ps1"
        ]
        shutdown = {
          command = "shutdown /s /t 10 /f /d p:4:1 /c Packer_Provisioning_Shutdown",
          timeout = "10s"
        }
      }
    }
    server = {
      "2012_R2" = {
      scripts = [
        # "${local.script_dir}/${var.os_family}/base_setup.ps1",
        "${local.script_dir}/${var.os_family}/provision.ps1",
        "${local.script_dir}/${var.os_family}/disable-windows-updates.ps1",
        "${local.script_dir}/${var.os_family}/disable-windows-defender.ps1",
        "${local.script_dir}/${var.os_family}/remove-one-drive.ps1",
        # "${local.script_dir}/windows/remove-apps.ps1",
        "${local.script_dir}/${var.os_family}/virtualbox-prevent-vboxsrv-resolution-delay.ps1",
        "${local.script_dir}/${var.os_family}/provision-winrm.ps1",
        "${local.script_dir}/${var.os_family}/enable-remote-desktop.ps1",
        "${local.script_dir}/${var.os_family}/eject-media.ps1"
      ]
    }
  }
}
}