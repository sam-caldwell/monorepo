#!/bin/bash -e

echo "reset the filesystem"
sudo umount /opt/data || true
#echo "y" | sudo mkfs.ext4 /dev/sdb
#sudo mkfs.btrfs /dev/sdb -f
sudo mkfs.btrfs -R free-space-tree -f -L data /dev/sdb
#sudo mount -o noatime,rw,ssd,compress=zlib:9,autodefrag, commit=120,noacl,nobarrier,discard=async /dev/sdb /opt/data
sudo mount -o noatime,rw,ssd,compress=lzo,autodefrag,commit=240,noacl,nobarrier,nodiscard /dev/sdb /opt/data
sudo chown $(whoami) /opt/data
echo 2048 | sudo tee /sys/block/sdb/queue/nr_requests

echo "start test"
go run main.go
