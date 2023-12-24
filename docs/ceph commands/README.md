## Cheat sheet for Ceph commands...

### General commands:
1. List all ceph nodes: `ceph node ls all`

### Cluster Map Display commands:
1. Display the monitor map: `ceph mon dump`
2. Display the manager map: `ceph mgr dump`
3. Display the OSD (Object Storage Daemon) map: `ceph osd dump`
4. Display the MSD (Metadata Storage Daemon) map: `ceph fs dump`
5. Display the CRUSH map: 
		`ceph osd getcrushmap -o {raw-filename}`.  
		(This will create a compiled crush map file which must be decompiled)
		`crushtool -d {raw-filename} -o {decomp-filename}`
		(This resulting `{decomp-filename}` can be displayed using `cat` 
		or any other similar tool.)
