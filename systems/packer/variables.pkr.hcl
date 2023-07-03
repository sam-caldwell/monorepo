#
# systems/packer/variables.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#

/*
 * An Operating System family is the general class of
 * operating system (e.g. windows, linux, bsd, darwin)
 */
variable "os_family" {
  description = "Operating system family"
  type        = string
  default     = "windows"
}
/*
 * A Specific operating system (e.g. windows,
 * ubuntu, centos, fedora, macos)
 */
variable "opsys" {
  description = "Operating system selector"
  type        = string
  default     = "windows"
}
/*
 * The operating system version (depends heavily on the
 * operating system, but for windows it could be 10, 11, etc)
 */
variable "os_version" {
  description = "Operating system version selector"
  type        = string
  default     = "10"
}
/*
 * The 'prefix' is the string we use at the start of a vm name
 */
variable "prefix" {
  description = "Prefix for vm name"
  type        = string
  default     = "vm_"
}
/*
 * Source is the hypervisor we will use to create our
 * virtual machine, such as parallels, vmware, virtualbox,
 * qemu, etc.
 */
variable "source" {
  description = "The builder source we will use (e.g. parallels)"
  type        = string
  default     = "parallels"
}
