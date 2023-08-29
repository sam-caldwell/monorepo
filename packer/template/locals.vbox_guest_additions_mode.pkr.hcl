locals {

  vbox_guest_additions_mode = var.vbox_guest_additions_mode == null ? local.default_vbox_guest_additions_mode : var.vbox_guest_additions_mode

}