#!/bin/sh -eux
#
#ubuntu_version="$(lsb_release -r | awk '{print $2}')";
#major_version="$(echo "$ubuntu_version" | awk -F. '{print $1}')";
#
#if [ "$major_version" -ge "18" ]; then
#echo "Create netplan config for enp0s5"
#cat <<EOF >/etc/netplan/00-installer-config.yaml;
#network:
#  version: 2
#  ethernets:
#    enp0s5:
#      dhcp4: true
#EOF
#else
#  # Adding a 2 sec delay to the interface up, to make the dhclient happy
#  echo "pre-up sleep 2" >> /etc/network/interfaces;
#fi
#
# Disable Predictable Network Interface names and use enp0s5
#[ -e /etc/network/interfaces ] && sed -i 's/en[[:alnum:]]*/enp0s5/g' /etc/network/interfaces;
#sed -i 's/GRUB_CMDLINE_LINUX="\(.*\)"/GRUB_CMDLINE_LINUX="net.ifnames=0 biosdevname=0 \1"/g' /etc/default/grub;
#update-grub;
