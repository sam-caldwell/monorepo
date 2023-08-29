locals {

  vbox_gfx_controller = var.vbox_gfx_controller == null ? local.default_windows_vbox_gfx_controller : var.vbox_gfx_controller

}