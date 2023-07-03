locals {
  linux_execute_command="echo 'vagrant' | {{ .Vars }} sudo -S -E sh -eux '{{ .Path }}'"
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
          "HOME_DIR=/home/vagrant", //ToDo: fix this
#          "http_proxy=${var.http_proxy}", //ToDo: fix this
#          "https_proxy=${var.https_proxy}", //ToDo: fix this
#          "no_proxy=${var.no_proxy}" //ToDo: fix this
        ]
        execute_command  = local.linux_execute_command
        iso              = ""
        iso_checksum     = ""
        boot_command     = [
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
          "hostname=vagrant ",
          "fb=false debconf/frontend=noninteractive ",
          "keyboard-configuration/modelcode=SKIP ",
          "keyboard-configuration/layout=USA ",
          "keyboard-configuration/variant=USA console-setup/ask_detect=false ",
          "passwd/user-fullname=__USER_NAME__ ", //ToDo: fix this
          "passwd/user-password=__USER_NAME__ ", //ToDo: fix this
          "passwd/user-password-again=__USER_PASSWORD__ ", //ToDo: fix this
          "passwd/username=__USER_NAME__ ", //ToDo: fix this
          "initrd=/install/initrd.gz -- <enter>"
        ]
      }
    }
  }
}