#
# systems/packer/windows10/variables.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
variable "iso_url" {
  description = "URL to the Windows 10 ISO"
  type        = string
}

variable "iso_checksum" {
  description = "Checksum of the Windows 10 ISO"
  type        = string
}

variable "iso_checksum_type" {
  description = "Checksum type of the Windows 10 ISO"
  type        = string
  default     = "sha256"
}

variable "disk_size"{
  description = "Disk size"
  type = string
  default = 20000
}

variable "username" {
  description = "Vagrant username"
  type        = string
}

variable "username" {
  description = "vagrant password"
  type        = string
}

