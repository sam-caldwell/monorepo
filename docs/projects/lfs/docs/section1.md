Create Partitions
=================

### Root Partition
```shell

dd if=/dev/urandom of=root.bin bs=1048576 count=20480
dd if=/dev/urandom of=boot.efi.bin bs=1024 count=500
dd if=/dev/urandom of=home.bin bs=1024 count=500
dd if=/dev/urandom of=usr.bin bs=1024 count=500
dd if=/dev/urandom of=usr.src.bin bs=1024 count=500
dd if=/dev/urandom of=var.bin bs=1024 count=500
dd if=/dev/urandom of=var.log.bin bs=1024 count=500
dd if=/dev/urandom of=var.lib.bin bs=1024 count=500
dd if=/dev/urandom of=boot.bin bs=1024 count=500
dd if=/dev/urandom of=boot.bin bs=1024 count=500
dd if=/dev/urandom of=boot.bin bs=1024 count=500
dd if=/dev/urandom of=boot.bin bs=1024 count=500
dd if=/dev/urandom of=boot.bin bs=1024 count=500
dd if=/dev/urandom of=boot.bin bs=1024 count=500
dd if=/dev/urandom of=boot.bin bs=1024 count=500
```
