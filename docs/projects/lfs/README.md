Linux From Scratch
==================

## Filesystem

### Master Image File

* We create the master image file (master.disk). This
  image file is greater than or equal to the total size
  of all partition image files (below).
* This image file will ultimately contain the entire
  installation.
* We can compress the master image file to save space.
* Ultimately the Master Image file can be written to disk,
  using `dd`.

### Partition Image Files

| File Name     | Size  | Partition | Format | Notes                                        |
|---------------|-------|-----------|--------|----------------------------------------------|
| boot.disk     | 500MB | /boot     | ext4   | readonly                                     |
| boot.efi.disk | 1MB   | /boot/efi |        | UEFI Partition                               |
| home.disk     | 20GB  | /home     | ext4   | noexec, nosuid, nodev                        |
| opt.disk      | 1GB   | /opt      | ext4   | nodev                                        |
| root.disk     | 20GB  | /         | ext4   | root partition                               |                          
| srv.disk      | 1MB   | /srv      | ext4   | nodev                                        |
| swap.dat      | 1GB   | /swap.dat | swap   | swap file                                    |                      
| tmp.disk      | 5GB   | /tmp      | tmpfs  | noexec, nosuid, nodev                        |
| usr.disk      | 20GB  | /usr      | ext4   | nodev                                        |
| usr.bin.disk  | 1GB   | /usr/bin  | ext4   | readonly, nodev                              |
| usr.src.disk  | 5GB   | /usr/src  | ext4   | nosuid, nodev                                |
| var.disk      | 20GB  | /var/     | ext4   | nosuid, nodev (note docker will use devices) |
| var.log.disk  | 5GB   | /var/log  | ext4   | noexec, nosuid, nodev                        |
| var.tmp.disk  | 1GB   | /var/tmp  | tmpfs  | noexec, nosuid, nodev                        |

* Each partition file can be created inside the master image file
  or copied into the master image file if created separately, but
  this will take up more disk space.
* Once we have the Master Image File mounted, we can `mount` and
  `chroot` into this environment.
* Using the mounted / chroot'd environment, Linux can be installed.

### Install Linux
