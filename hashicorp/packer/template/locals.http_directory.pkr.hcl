locals {
  http_directory = var.http_directory == null ? "${path.root}/http" : var.http_directory
}