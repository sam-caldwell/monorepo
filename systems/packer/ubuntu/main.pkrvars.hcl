#
# systems/packer/ubuntu/main.pkrvars.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
iso_url           = "create.sh/iso/ubuntu-22.04.1-live-server-amd64.iso"
iso_checksum      = "10f19c5b2b8d6db711582e0e27f5116296c34fe4b313ba45f9b201a5007056cb"
iso_checksum_type = "sha256"

disk_size = 20000

username = "vagrant"
password = "vagrant"