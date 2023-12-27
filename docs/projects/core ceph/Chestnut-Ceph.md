## Objectives
1. Construct a `ceph` cluster using Rock64, Raspberry Pi or NUC devices in a set of two-post racks, where each SOC compute device will attach to at least 1-5TB of storage and provide s3-compatible object storage.
2. Create a method of operating the `ceph` cluster using POE power such that individual devices can be controlled programmatically through the POE switch(es).

## Assumptions
1. Capacity:
	1. A two-post Rack can contain 4-5 SOC devices per 1 u row.
	2. A two-post Rack can contain at least 1 x 1 TB SSD/HDD storage device per SOC in the same 1 u, connected via USB or SATA-III.  (Note: SATA-III is preferred for throughput).
	3. A standard 45 u two-post rack, assuming 5 u at the bottom open for ventilation, would have 40 u of capacity @ 5 SOC-disk pairs per row or 200 devices.
	4. Thus each two-post rack would support 200 TB of storage.
2. Networking
	1. Assume each SOC has 1 x 1 Gbps Ethernet.
	2. This means we would need 200 Ethernet ports per rack.
	3. Assume 40 ports per switch, we would need 5 x 40 port 1 Gbps switches with up-link and down-link to create 200 ports of capacity per rack.
	4. Assume each switch has redundant power supplies.  This means 10 x 110 VAC power outlets on two 20 A circuits.
	5. Assume two UPS battery devices with 3000 W capacity per device with 5 outlets per device.  This reduces our physical 20 A connections to four outlets minimum, one per circuit.
3. Power
	1. Assuming 3 A, 5 V, each device would draw 15 W.
	2. Assuming 200 devices per rack, this would mean 3000 W per rack of power.
	3. See notes on power for network devices which will power the SOC racks.
4. Distribution/HA
	1. We will mitigate risk by creating three separate racks (for a total 200 TB with 3 replicas).
	2. Initial locations
		1. Rack 00 will be located in the network closet
		2. Rack 01 will be located in the office closet
		3. Rack 02 will be located in the basement
	3. Future locations
		1. Rack 03 will be located in the network closet
		2. Rack 04 will be located in the office closet
		3. Rack 05 will be located in the basement.
5. Performance considerations
	1. Higher latency (object store) storage should be happy with 3 replicas using lower-performing devices.  This should work well with the Raspberry PI/Rock64 devices.
	2. Lower Latency (CephFS/Block Storage) will require higher performing disks and more CPU/RAM per unit (requiring NUC or similar devices).
	3. Question: Can we improve performance by creating multiple replicas >3 to distribute load on lower-performing devices at a lower per GB-hour price point?