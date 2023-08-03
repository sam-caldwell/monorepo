packer {
  required_version = ">= 1.9.2"
  required_plugins {
    parallels = {
      version = ">= 1.1.0"
      source  = "github.com/parallels/parallels"
    }
    qemu = {
      version = ">= 1.0.9"
      source  = "github.com/hashicorp/qemu"
    }
    vagrant = {
      version = ">= 1.0.3"
      source  = "github.com/hashicorp/vagrant"
    }
    virtualbox = {
      version = ">= 0.0.5"
      source  = "github.com/hashicorp/virtualbox"
    }
    vmware = {
      version = ">= 1.0.10"
      source  = "github.com/Stromweld/vmware"
    }
    # Temp switch till bug fix for x86 tools location is fixed
    # vmware = {
    #   version = ">= 1.0.8"
    #   source  = "github.com/hashicorp/vmware"
    # }
    windows-update = {
      version = ">= 0.14.3"
      source  = "github.com/rgl/windows-update"
    }
  }
}