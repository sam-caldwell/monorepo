locals {
  /*
   * General
   */
  iso_checksum_type = "sha256"
  general = {
    disk = {
      size = 20000
    }
    user              = {
      username = "vagrant"
      password = "vagrant"
    }
  }
  linux = {
    ubuntu = {
      iso=""
      iso_checksum=""
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
        "hostname=vagrant ",
        "fb=false debconf/frontend=noninteractive ",
        "keyboard-configuration/modelcode=SKIP ",
        "keyboard-configuration/layout=USA ",
        "keyboard-configuration/variant=USA console-setup/ask_detect=false ",
        "passwd/user-fullname=__USER_NAME__ ",
        "passwd/user-password=__USER_NAME__ ",
        "passwd/user-password-again=__USER_PASSWORD__ ",
        "passwd/username=__USER_NAME__ ",
        "initrd=/install/initrd.gz -- <enter>"
      ]
    }
    centos = {

    }
    fedora = {

    }
  }
  windows = {
    10 = {
      boot_command=[]
      is_windows    = true
      os_name       = "windows"
      os_version    = "10"
      os_arch       = "x86_64"
      iso           = "https://create.sh/iso/windows.10.iso"
      iso_checksum      = "ef7312733a9f5d7d51cfa04ac497671995674ca5e1058d5164d6028f0938d668"
      guest_os_type = {
        parallels = "win-10"
        vbox      = "Windows10_64"
        vmware    = "windows9srv-64"
      }
      communicator = "winrm"
      shutdown = "shutdown /s /t 10 /f /d p:4:1 /c Packer_Provisioning_Shutdown",
    }
    11 = {
      boot_command=[]
      is_windows           = true
      os_name              = "windows"
      os_version           = "11"
      os_arch              = "x86_64"
      iso                  = "https://create.sh/iso/windows.11.iso"
      iso_checksum             = "ebbc79106715f44f5020f77bd90721b17c5a877cbc15a3535b99155493a1bb3f"
      vmware_guest_os_type = "windows9srv-64"
      shutdown             = "shutdown /s /t 10 /f /d p:4:1 /c Packer_Provisioning_Shutdown",
    }
  }


}