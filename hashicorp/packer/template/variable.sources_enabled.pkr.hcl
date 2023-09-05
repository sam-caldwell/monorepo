variable "sources_enabled" {
  type    = list(string)
  default = [
    "source.parallels-iso.vm",
#    "source.qemu.vm",
#    "source.virtualbox-iso.vm",
#    "source.vmware-iso.vm",
  ]
  description = "Build Sources to use for building vagrant boxes"
}