#
# systems/packer/windows10/main.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
build {
  sources = [
    "sources.parallels-iso.vm",
    #    "sources.virtualbox-iso.vm",
    #    "sources.vmware-iso.vm",
    #        "sources.qemu.vm"
  ]

  provisioner "powershell" {
    inline = [
      "choco install vagrant -y",
      "Restart-Computer -Force"
    ]
  }

  post-processor "vagrant" {
    output = join("", ["build/", "{{.Provider}}", ".", var.opsys, ".", var.os_version, ".box"])
  }
}
