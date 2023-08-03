variable "dns" {
  description = "This is the list of dns servers which our virtual machine will use."
  type    = list(string)
  default = [
    "1.1.1.1",
    "8.8.8.8"
  ]
}
