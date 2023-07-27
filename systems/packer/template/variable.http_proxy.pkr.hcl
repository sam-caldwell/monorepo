variable "http_proxy" {
  type        = string
  default     = env("http_proxy")
  description = "Http proxy url to connect to the internet"
}