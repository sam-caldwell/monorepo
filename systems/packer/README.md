Packer Builder
==============

## Description

This project contains hashicorp packer build manifests
which will generate vagrant boxes and other virtual
machine images from a uniform configuration-as-code source.

These vagrant boxes and virtual machines are intended as--

1. Base immutable machines
2. Test machines for automated testing
3. Cloud-deployable machines
4. Virtual hosts for deployment to local hosting

## Helpful resources
### Windows
  * [virtio drivers](https://docs.fedoraproject.org/en-US/quick-docs/creating-windows-virtual-machines-using-virtio-drivers/index.html#virtio-win-direct-downloads)
  * [virtio-win](https://fedorapeople.org/groups/virt/virtio-win/direct-downloads/archive-virtio/?C=M;O=D)
  * [Windows: LogonCount known issue](https://learn.microsoft.com/en-us/windows-hardware/customize/desktop/unattend/microsoft-windows-shell-setup-autologon-logoncount#logoncount-known-issue)
  * [Windows: Enable-PSRemoting](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.core/enable-psremoting?view=powershell-7.3)
  * [How to prevent windows 10 prompting for network location](https://social.technet.microsoft.com/Forums/en-US/42a12344-21e9-4f14-946f-6b4743b20403/how-to-prevent-windows-10-osd-prompting-for-network-discovery?forum=configmanagerosd)

## Usage
### To Validate/Lint Manifests
`make validate/packer`

### To Build an image
`make build/packer`

