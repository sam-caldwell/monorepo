#!/bin/bash -e

echo "reset the filesystem"
sudo umount /opt/data || true
#echo "y" | sudo mkfs.ext4 /dev/sdb
sudo mkfs.btrfs /dev/sdb -f
sudo mount -o noatime,rw,ssd,compress=zlib,autodefrag,nodatacow,commit=60 /dev/sdb /opt/data
sudo chown $(whoami) /opt/data
echo "start test"
go run main.go
