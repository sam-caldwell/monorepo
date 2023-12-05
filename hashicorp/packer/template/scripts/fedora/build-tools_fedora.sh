#!/bin/bash -eux
# Installing build sqldbtest here because Fedora 22+ will not do so during kickstart
dnf -y install kernel-headers kernel-devel-"$(uname -r)" elfutils-libelf-devel gcc make perl
