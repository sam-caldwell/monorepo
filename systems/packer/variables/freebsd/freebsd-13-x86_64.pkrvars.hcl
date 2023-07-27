os_name                 = "freebsd"
os_version              = "13.2"
os_arch                 = "x86_64"
iso_url                 = "https://download.freebsd.org/releases/amd64/amd64/ISO-IMAGES/13.2/FreeBSD-13.2-RELEASE-amd64-disc1.iso"
iso_checksum            = "file:https://download.freebsd.org/releases/amd64/amd64/ISO-IMAGES/13.2/CHECKSUM.SHA256-FreeBSD-13.2-RELEASE-amd64"
parallels_guest_os_type = "freebsd"
vbox_guest_os_type      = "FreeBSD_64"
vmware_guest_os_type    = "freebsd-64"
boot_command            = ["<wait><esc><wait>boot -s<wait><enter><wait><wait10><wait10>/bin/sh<enter><wait>mdmfs -s 100m md1 /tmp<enter><wait>mdmfs -s 100m md2 /mnt<enter><wait>dhclient -p /tmp/dhclient.em0.pid -l /tmp/dhclient.lease.em0 em0<enter><wait><wait5>fetch -o /tmp/installerconfig http://{{ .HTTPIP }}:{{ .HTTPPort }}/freebsd/installerconfig && bsdinstall script /tmp/installerconfig<enter><wait>"]
