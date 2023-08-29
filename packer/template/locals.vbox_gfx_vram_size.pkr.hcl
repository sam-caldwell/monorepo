locals {

  vbox_gfx_vram_size        = var.vbox_gfx_controller == null ? local.default_vbox_gfx_vram_size : var.vbox_gfx_vram_size

}