variable "no_proxy" {
  type        = string
  default     = env("no_proxy")
  description = "No Proxy"
}