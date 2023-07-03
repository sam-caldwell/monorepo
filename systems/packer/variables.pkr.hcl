#
# systems/packer/variables.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#

variable "os_family" {
  description = "Operating system family"
  type        = string
  //Windows, linux, darwin
}

variable "opsys" {
  description = "Operating system selector"
  type        = string
  //windows, ubuntu, centos, fedora, macos
}

variable "os_version" {
  description = "Operating system version selector"
  type        = string
  //10, 22.04...
}

variable "prefix" {
  description = "Prefix for vm name"
  type        = string
  default     = "vm_"
}