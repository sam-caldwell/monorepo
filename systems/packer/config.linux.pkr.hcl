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
        ]
      }
    }
  }
}