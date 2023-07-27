vm_name                 = "ubuntu-server"
os_name                 = "ubuntu"
os_version              = "22.04"
os_arch                 = "x86_64"
iso_url                 = "https://releases.ubuntu.com/jammy/ubuntu-22.04.2-live-server-amd64.iso"
iso_checksum            = "5e38b55d57d94ff029719342357325ed3bda38fa80054f9330dc789cd2d43931"
hyperv_generation       = 2
parallels_guest_os_type = "ubuntu"
vbox_guest_os_type      = "Ubuntu_64"
vmware_guest_os_type    = "ubuntu-64"
boot_command            = [
  "<wait>c<wait>",
  "set gfxpayload=keep<enter><wait>",
  "linux /casper/vmlinuz quiet autoinstall ds=nocloud-net\\;s=http://{{.HTTPIP}}:{{.HTTPPort}}/ubuntu/ ---<enter>",
  "<wait>",
  "initrd /casper/initrd<wait>",
  "<enter><wait>boot<enter>",
  "<wait15><enter>"
]
