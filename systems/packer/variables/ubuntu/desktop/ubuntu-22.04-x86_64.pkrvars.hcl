vm_name                 = "ubuntu-server"
os_name                 = "ubuntu"
os_version              = "22.04"
os_arch                 = "x86_64"
iso_url                 = "https://releases.ubuntu.com/22.04.2/ubuntu-22.04.2-desktop-amd64.iso"
iso_checksum            = "b98dac940a82b110e6265ca78d1320f1f7103861e922aa1a54e4202686e9bbd3"
hyperv_generation       = 2
parallels_guest_os_type = "ubuntu"
vbox_guest_os_type      = "Ubuntu_64"
vmware_guest_os_type    = "ubuntu-64"
boot_command            = [
  "<wait>c<wait>",
  "set gfxpayload=keep<enter><wait>",
  "linux /casper/vmlinuz quiet autoinstall ds=nocloud-net\\;s=http://{{.HTTPIP}}:{{.HTTPPort}}/ubuntu/ ---<enter>",
  "<wait>initrd /casper/initrd<wait><enter><wait>boot<enter>",
  "<wait>"
]
