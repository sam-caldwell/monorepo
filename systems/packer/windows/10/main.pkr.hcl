#
# systems/packer/windows10/main.pkr.hcl
# (c) Sam Caldwell.  See LICENSE.txt
#

variable "iso_url" {
  description = "URL to the Windows 10 ISO"
  type        = string
}

variable "iso_checksum" {
  description = "Checksum of the Windows 10 ISO"
  type        = string
}

variable "iso_checksum_type" {
  description = "Checksum type of the Windows 10 ISO"
  type        = string
  default     = "sha256"
}

variable "disk_size"{
  description = "Disk size"
  type = string
  default = 20000
}

variable "username" {
  description = "Vagrant username"
  type        = string
}

variable "password" {
  description = "vagrant password"
  type        = string
}


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

source "parallels-iso" "windows10" {
  vm_name = "windows10"

  parallels_tools_flavor="win"
  shutdown_command       = "echo 'vagrant'|sudo -S shutdown -P now"
  iso_urls = [ var.iso_url ]
  iso_checksum = var.iso_checksum
#  iso_checksum_type = var.iso_checksum_type

  guest_os_type = "win-10"

  disk_size = var.disk_size

  communicator = "winrm"
  winrm_username = var.username
  winrm_password = var.password
}
build {
  sources = [
    "sources.parallels-iso.windows10",
    #    "sources.virtualbox-iso.windows10",
    #    "sources.vmware-iso.windows10",
    #    "sources.qemu.windows10"
  ]

  provisioner "powershell" {
    inline = [
      "choco install vagrant -y",
      "Restart-Computer -Force"
    ]
  }

  post-processor "vagrant" {
    output = "build/windows_10_{{.Provider}}.box"
  }
}
