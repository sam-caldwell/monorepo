locals {

  default_windows_vbox_gfx_controller = var.is_windows ? "vboxsvga" : "vmsvga"
}