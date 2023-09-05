variable "iso_checksum" {
  type        = string
  default     = "file:https://releases.ubuntu.com/jammy/SHA256SUMS"
  description = "ISO download checksum"
}