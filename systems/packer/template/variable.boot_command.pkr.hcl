variable "boot_command" {
  type        = list(string)
  default     = []
  description = "Commands to pass to gui session to initiate automated install"
}