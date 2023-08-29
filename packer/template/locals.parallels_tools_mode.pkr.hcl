locals {

  parallels_tools_mode = var.parallels_tools_mode == null ? (

  var.is_windows ? "attach" : "upload"

  ) : var.parallels_tools_mode

}