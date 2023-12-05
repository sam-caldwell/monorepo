#!/bin/sh -eux

case "$PACKER_BUILDER_TYPE" in
vmware-iso|vmware-vmx)
    echo "install open-vm-tools"
    apt-get install -y open-vm-sqldbtest;
    mkdir /mnt/hgfs;
    systemctl enable open-vm-sqldbtest
    systemctl start open-vm-sqldbtest
    echo "platform specific vmware.sh executed";
esac
