#
# systems/packer/packer.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#
packer {
  required_plugins {
    amazon = {
      version = ">= 1.2.6"
      source = "github.com/hashicorp/amazon"
    }
    digitalocean = {
      version = ">= 1.0.4"
      source  = "github.com/digitalocean/digitalocean"
    }
    docker = {
      version = ">= 0.0.7"
      source  = "github.com/hashicorp/docker"
    }
    parallels = {
      version = ">= 1.0.0"
      source  = "github.com/hashicorp/parallels"
    }
    qemu = {
      version = ">= 1.0.0"
      source  = "github.com/hashicorp/qemu"
    }
    virtualbox = {
      version = ">= 0.0.1"
      source = "github.com/hashicorp/virtualbox"
    }
    vmware = {
      version = ">= 1.0.0"
      source = "github.com/hashicorp/vmware"
    }
  }
}
