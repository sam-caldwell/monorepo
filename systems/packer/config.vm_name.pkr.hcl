/*
 * systems/packer/config.vm_name.pkr.hcl
 * (c) Sam Caldwell.  See LICENSE.txt
 *
 * This file calculates our vm_name.
 */
locals {
  vm_name = join("", [var.prefix, var.opsys, var.os_version])
}