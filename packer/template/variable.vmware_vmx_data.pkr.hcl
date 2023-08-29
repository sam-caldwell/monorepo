variable "vmware_vmx_data" {
  type    = map(string)
  default = {
    "cpuid.coresPerSocket"    = "2"
    "ethernet0.pciSlotNumber" = "32"
    "svga.autodetect"         = true
    "usb_xhci.present"        = true
  }
}