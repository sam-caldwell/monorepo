#
# systems/packer/windows10/main.pkrvars.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
iso_url           = "https://create.sh/iso/windows.10.iso"
iso_checksum      = "ef7312733a9f5d7d51cfa04ac497671995674ca5e1058d5164d6028f0938d668"
iso_checksum_type = "sha256"

disk_size = 20000

username = "vagrant"
password = "vagrant"