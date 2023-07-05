/*
 * systems/packer/config.linux.pkr.hcl
 * (c) Sam Caldwell.  See LICENSE.txt
 *
 * Windows-specific configuration for windows 10.
 */
locals {
  linux_execute_command = "echo 'vagrant' | {{ .Vars }} sudo -S -E sh -eux '{{ .Path }}'"
}
locals {
  /*
   * Linux-specific configuration
   */
  //os_family
  linux = {
    //opsys
    ubuntu = {
      //os_version
      22.04 = {
        environment_vars = [
          "HOME_DIR=/home/${local.fact.user.username}",
          "http_proxy=${local.fact.http_proxy}",
          "https_proxy=${local.fact.https_proxy}",
          "no_proxy=${local.fact.no_proxy}"
        ]
        execute_command = local.linux_execute_command
        iso             = ""
        iso_checksum    = ""
        boot_command = [
          "<enter><wait><f6><wait><esc><wait>",
          "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
          "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
          "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
          "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
          "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
          "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
          "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
          "<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
          "<bs><bs><bs>",
          "/install/vmlinuz noapic ",
          "file=/floppy/preseed.cfg ",
          "debian-installer=en_US auto locale=en_US kbd-chooser/method=us ",
          "hostname=${local.fact.hostname} ",
          "fb=false debconf/frontend=noninteractive ",
          "keyboard-configuration/modelcode=SKIP ",
          "keyboard-configuration/layout=USA ",
          "keyboard-configuration/variant=USA console-setup/ask_detect=false ",
          "passwd/user-fullname=${local.fact.user.username} ",
          "passwd/user-password=${local.fact.user.password} ",
          "passwd/user-password-again=${local.fact.user.password} ",
          "passwd/username=${local.fact.user.username} ",
          "initrd=/install/initrd.gz -- <enter>"
        ]
        scripts = [
          "${local.script_dir}/${var.opsys}/update_${var.opsys}.sh",
          "${local.script_dir}/_common/motd.sh",
          "${local.script_dir}/_common/sshd.sh",
          "${local.script_dir}/${var.opsys}/networking_${var.opsys}.sh",
          "${local.script_dir}/${var.opsys}/sudoers_${var.opsys}.sh",
          "${local.script_dir}/_common/vagrant.sh",
          "${local.script_dir}/${var.opsys}/systemd_${var.opsys}.sh",
          "${local.script_dir}/_common/virtualbox.sh",
          "${local.script_dir}/_common/vmware_debian_ubuntu.sh",
          "${local.script_dir}/_common/parallels.sh",
          "${local.script_dir}/${var.opsys}/hyperv_${var.opsys}.sh",
          "${local.script_dir}/${var.opsys}/cleanup_${var.opsys}.sh",
          "${local.script_dir}/_common/minimize.sh"
        ]
      }
    }
  }
}