build {
  sources = [
#    "sources.virtualbox-iso.windows10",
#    "sources.vmware-iso.windows10",
    "sources.parallels-iso.windows10",
#    "sources.qemu.windows10"
  ]

  provisioner "powershell" {
    inline = [
      "choco install vagrant -y",
      "Restart-Computer -Force"
    ]
  }

  post-processor "vagrant" {
    output = "windows_10_{{.Provider}}.box"
  }
}
