variable "https_proxy" {
  type        = string
  default     = env("https_proxy")
  description = "Https proxy url to connect to the internet"
}