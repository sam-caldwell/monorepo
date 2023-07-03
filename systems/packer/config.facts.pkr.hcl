/*
 * systems/packer/config.facts.pkr.hcl
 * (c) Sam Caldwell.  See LICENSE.txt
 *
 * This file establishes certain facts based on vars and locals
 */
locals {
  fact = {
    source = {
      enabled = {
        hyperv     = (var.source=="hyperv")?"sources.hyperv-iso.vm" : ""
        parallels  = (var.source=="parallels")?"sources.parallels-iso.vm" : ""
        qemu       = (var.source=="qemu")?"sources.qemu.vm" : ""
        virtualbox = (var.source=="virtualbox")?"sources.virtualbox-iso.vm" : ""
        vmware     = (var.source=="vmware")?"sources.vmware-iso.vm" : ""
      }
    }
    is_windows     = var.os_family=="windows"
  }
}