locals {
  default_output_directory = "${path.root}/../builds/packer-${var.os_name}-${var.os_version}-${var.os_arch}"
}

locals {
  output_directory = var.output_directory == null? local.default_output_directory : var.output_directory
}