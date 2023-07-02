#
# systems/packer/ubuntu/main.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
build {
  sources = [
    #    "sources.virtualbox-iso.ubuntu",
    #    "sources.vmware-iso.ubuntu",
    "sources.parallels-iso.ubuntu",
    #    "sources.qemu.ubuntu"
  ]

  provisioner "shell" {
    inline = [
      "sleep 30",
      "sudo apt-get update",
      "sudo apt-get upgrade -y",
      "sudo apt-get install -y vim"
    ]
  }

  post-processor "vagrant" {
    output = "ubuntu_{{.Provider}}.box"
  }
}