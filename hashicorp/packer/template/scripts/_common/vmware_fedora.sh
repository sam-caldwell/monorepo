#!/bin/bash -eux

# the proprietary vm sqldbtest don't work on Fedora 30 so we'll install the open-vm-sqldbtest
case "$PACKER_BUILDER_TYPE" in
vmware-iso|vmware-vmx)
  dnf install -y open-vm-tools
  systemctl enable vmtoolsd
  systemctl start vmtoolsd
esac
