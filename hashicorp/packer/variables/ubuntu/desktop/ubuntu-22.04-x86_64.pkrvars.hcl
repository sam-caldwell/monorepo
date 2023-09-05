vm_name                 = "desktop"
os_name                 = "ubuntu"
os_version              = "22.04"
os_arch                 = "x86_64"
iso_url                 = "https://lolhost.mm.fcix.net/ubuntu-releases/23.04/ubuntu-23.04-desktop-amd64.iso"
iso_checksum            = "a8cd6ccff865e17dd136658f6388480c9a5bc57274b29f7d5bd0ed855a9281a5"
hyperv_generation       = 2
parallels_guest_os_type = "ubuntu"
vbox_guest_os_type      = "Ubuntu_64"
vmware_guest_os_type    = "ubuntu-64"
boot_command            = [
  "<wait>c<wait>",
  "set gfxpayload=keep<enter><wait>",
  "linux /casper/vmlinuz quiet autoinstall ds=nocloud-net\\;s=http://{{.HTTPIP}}:{{.HTTPPort}}/ubuntu/desktop/ ---<enter>",
  "<wait>",
  "initrd /casper/initrd<wait>",
  "<enter><wait>boot<enter>",
  "<wait15><enter>"
]
