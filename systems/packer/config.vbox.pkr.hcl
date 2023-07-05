locals {

  // Resolve our virtual box configuration...
  #  vbox= {
  #    /*
  #   * VIRTUALBOX: vbox_gfx_controller
  #   */
  #    vbox_gfx_controller = var.vbox_gfx_controller == null ? (
  #    local.fact.is_windows ? "vboxsvga" : "vmsvga"
  #    ) : var.vbox_gfx_controller
  #    /*
  #   * VIRTUALBOX: vbox_gfx_vram_size
  #   */
  #    vbox_gfx_vram_size = var.vbox_gfx_controller == null ? (
  #    local.fact.is_windows ? 128 : 33
  #    ) : var.vbox_gfx_vram_size
  #    /*
  #   * VIRTUALBOX: vbox_guest_additions_mode
  #   */
  #    vbox_guest_additions_mode = var.vbox_guest_additions_mode == null ? (
  #    local.fact.is_windows && var.hyperv_generation == 1 ? "attach" : "upload"
  #    ) : var.vbox_guest_additions_mode
  #    /*
  #   * VIRTUALBOX: vbox_source
  #   */
  #    # virtualbox-ovf
  #    vbox_source = var.vbox_source == null ? (
  #    var.opsys == "amazonlinux" ? "${path.root}/amz_working_files/amazon2.ovf" : null
  #    ) : var.vbox_source
  #  }
}