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
