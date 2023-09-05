locals {
  linux_environment_variables = [
    "HOME_DIR=/home/vagrant",
    "http_proxy=${var.http_proxy}",
    "https_proxy=${var.https_proxy}",
    "no_proxy=${var.no_proxy}"
  ]
}