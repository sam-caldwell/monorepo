/*
 * systems/packer/packer.pkr.hcl
 * (c) Sam Caldwell.  See LICENSE.txt
 *
 * This file defines our packer version and plugins.
 * The contents of this file must be statically defined
 * due to limitations of packer.
 */
packer {
  required_version = ">= 1.9.0"
  required_plugins {
    amazon = {
      version = ">= 1.2.6"
      source  = "github.com/hashicorp/amazon"
    }
    digitalocean = {
      version = ">= 1.0.4"
      source  = "github.com/digitalocean/digitalocean"
    }
    docker = {
      version = ">= 0.0.7"
      source  = "github.com/hashicorp/docker"
    }
    hyperv = {
      version = ">= 1.0.0"
      source  = "github.com/hashicorp/hyperv"
    }
    parallels = {
      version = ">= 1.0.1"
      source  = "github.com/hashicorp/parallels"
    }
    qemu = {
      version = ">= 1.0.8"
      source  = "github.com/hashicorp/qemu"
    }
    vagrant = {
      version = ">= 1.0.2"
      source  = "github.com/hashicorp/vagrant"
    }
    virtualbox = {
      version = ">= 0.0.1"
      source  = "github.com/hashicorp/virtualbox"
    }
    vmware = {
      version = ">= 1.0.8"
      source  = "github.com/hashicorp/vmware"
    }
    windows-update = {
      version = ">= 0.14.1"
      source  = "github.com/rgl/windows-update"
    }
  }
}
