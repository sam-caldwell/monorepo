## Why use `Ceph`?
1. `Ceph` is a petabyte scale, flexible solution for providing--
	1. S3-compatible `Ceph Object Storage`
	2. Distributed Block Storage (similar to iSCSI-DRBD clusters of old)
	3. Distributed File Storage (`CephFS`), similar to `GlusterFS` or `AWS EFS`.
2. `Ceph` is designed from the ground up for HA and gradual scaling over time.
3. `Ceph` is open source
4. Hardware requirements are modest, and we can build a `Ceph` cluster across a fleet of heterogenous hardware--including Rock64, Raspberry PI and other devices.

---
## Architectural Concepts...

1. General Architecture and Storage Types: ![[Pasted image 20231224015840.png]]
2. `Ceph` deploys as a cluster of microservices (discussed later).
	1. Each "`Ceph` Node" is an independent unit of the larger `Ceph` cluster.
	2. Each Ceph Node communicates with the other nodes in the cluster to replicate and redistribute data or manage the cluster.
	3. some Ceph Nodes are data- or metadata-specific while others are responsible for cluster management.
3. At the core is RADOS:
	1. RADOS - Reliable Autonomic Distributed Object Store
	2. RADOS is the underlying object store that provides scalable service for variably sized objects.
	3. At the heart of RADOS is the "object" (an atomic unit of storage) identified by a fixed-length numeric identifier and some number of named attributes.  The object itself is a variable-sized payload of bytes.
	4. Objects are stored in "object pools" where each pool has a name that forms a distinct object namespace.
	5. These "Object pools" also have a set of parameters that determine how the object is stored (e.g. replication level and a mapping rule to describe how replicas should be distributed across the cluster).
	6. A RADOS cluster can consist of a large number of pools and servers.
	7. The "storage services" (called Object Storage Daemons, or "OSDs") represent the servers which are responsible for actual data storage.![[Pasted image 20231224020214.png]]
	8. ***A key feature of RADOS is that the OSDs are able to operate with relative autonomy when it comes to recovering from failures or migrating data in response to cluster expansion.***
4. A `Ceph` Cluster consists of four (4) microservices (many with multiple instances for HA):
	1. Ceph Monitor (`ceph-monitor`)
		1. Maintains a master copy of the "Cluster Map" which describes the state and running configuration of the Ceph cluster.
		2. The "cluster maps", include the--
			1. ***Monitor map***
				 * Contains cluster `fsid`, 
				 * Identifies the position, name, address and TCP port of each monitor instance.
				 * Specifies the map's creation and last modified time stamp (epoch).
			 2. ***Manager map***
				 1. Contains the cluster `fsid`,
				 2. Identifies the position, name, address and TCP port of each manager instance
				 3. Specifies the map's creation and last modified time stamp (epoch).
			 3. ***OSD Map***
				 1. Contains the cluster `fsid`,
				 2. Specifies the map's creation and last modified timestamp (epoch).
				 3. Lists the pools, replica sizes, PG numbers and status of each OSD instance.
			 4. ***MDS Map***
				 1. Contains the cluster `fsid`,
				 2. Specifies the map's creation and last modified timestamp (epoch).
				 3. Contains the pool for storing metadata (about objects), a list of metadata (MDS) servers, their connection information and status (e.g. up, in)
			 5. ***CRUSH Map***
				 1. Contains the cluster `fsid`,
				 2. Contains a list of storage devices.
				 3. Identifies the "Failure Domain Hierarchy" (device, host, rack, row, room) and rules for traversing the hierarchy when storing data to ensure fault tolerance.
			 6. ***PG Map*** (Placement Group)
				 1. Contains the cluster `fsid`.
				 2. Contains the PG version and timestamp
				 3. Contains the last OSD map epoch, the full ratios and details of each placement group, including--
					 1. PG ID
					 2. the 'up' set
					 3. the 'acting' set
					 4. the state of the PG (e.g. "active + clean")
					 5. data usage statistics for each pool
		4. The `ceph-monitor` also manages authentication between clients and the cluster microservices. 
		5. A cluster can run with only ONE `ceph-monitor` but it will be a SPOF:
			1. <span style="color:red">The cluster will not operate without these maps.</span>(The maps are critical to allow the various Ceph daemons to coordinate with one another. )
			2. At least three (3) `ceph-monitor` instances are recommended for HA.
	2. Ceph Manager (`ceph-mgr`)
		1. Responsible for keeping track of runtime metrics and the current state of the Ceph cluster, including storage utilization, current performance metrics and system load
		2. Provides Web UI to expose cluster state.
		4. Provides API.
		5. Requires two (2) instances for HA, but can run with only one (1) SPOF instance.
	3. Ceph Object Storage Daemon (`ceph-osd`)
		1. Stores object data
		2. Performs any replication, recovery and rebalancing of data (see CRUSH).
		3. Provides some monitoring information to `ceph-monitor` and `ceph-mgr`.
		4. Checks other `ceph-osd` instances for their heartbeats.
		5. Requires at least three (3) OSDs for HA.
		6. Interacts with logical disks to store data.
	 4. Ceph Metadata Server (`ceph-mds`)
		 1. Required for `CephFS` (a POSIX-compliant  file system).
		 2. Tracks and stores metadata about the file system (and associated objects).
		 3. Very CPU intensive
	 5. Ceph Object Gateway (`ceph-rgw`)
		 1. Required for s3-compatible object storage services.
		 2. RGW-->RADOS Gateway
		 3. Requires one (1) at a minimum but three (3) are recommended for HA.
---
## Hardware Requirements
* Generally `ceph` is intended to run on commodity hardware ranging from Raspberry Pi / Rock64 to top-line servers.
* The cluster scales horizontally, making a thin layer of lighter hardware possible.
* ***CPU Notes:***
	* General
		* Hyper-threading is usually beneficial for `ceph` clusters.
	* For Monitor and Manager Nodes
		* Not CPU intensive.
		* Should have CPU limits and requirements configured to avoid contention.
	* For Metadata services (`ceph-mds`)
		* Require more CPU.  CPU Intensive.
		* Metadata service is single threaded, so multi-core CPUs have less value. 
			* It may be worth while to--
				* Run multiple instances of `ceph-mds` on a multi-core system.
				* Colocate `ceph-mds` with other CEPH microservices.
	* For OSD Nodes
		* Higher IOPS increases CPU load to--
			* Run RADOS
			* Calculate data placement (CRUSH) 
			* Replicate data
			* Maintain the local copies of the cluster map.
		* Exact CPU requirements vary depending on storage devices (NVMe, SSD, etc.)
		* IOPS per core (thread) is a good metric.
* ***RAM Notes:***
	* General
		*  Nodes require more memory during startup and rebalancing (especially for Monitor, Manager and OSD nodes).
	* For Monitor and Manager nodes--
		* Memory intensive (but CPU moderate)
		* RAM grows as the number of OSDs increases.
		* For clusters up to 300 OSDs, 64 GB is suggested.  Larger could benefit from 128GB.
	* For Metadata servers--
		* 1GB at minimum but increased memory benefits the cache size (`mds_cache_memory_limit`) and performance.
	* For OSD Nodes
		* 4DB per `BlueStore` OSD is default (`osd_memory_target`)
		* 8GB per `BlueStore` OSD is advised.
		* Where per OSD throughput is common (e.g. 256GB/OSD), or where there are many small objects per OSD, setting `osd_memory_target` to more than 8GB is recommended.  This is especially true for fast NVMe OSDs.
		* When provisioning, recovering,  rebalancing or starting more memory is needed.
		* A 20% memory margin is recommended because the kernel may lag the OSD application in reclaiming freed memory.
		* OSD may exceed the `osd_memory_target` value from time to time.  Having ***SOME*** swap in place to avoid OOM may be advisable (though it will harm performance).
			* The `ceph` project advises against swap.  But if swap can allow the kernel to overcome its lag, the performance drop may be acceptable.  Testing needed.
* ***Disk / Storage Notes***
	* General
		* Separate operating system and ceph partitions on different devices.
		* OSD data and BlueStore WAL+DB data on separate partitions.
		* Partition alignment can be a performance hurdle.
	* Monitor/Manager nodes
		* Disk intensive.  Recommend fast SSDs
	* Metadata servers
		* Disk intensive. 
		* Store metadata on a separate disk from OpSys
* ***Networking***
	* General
		* All services can run on single 1Gbps wired links.
		* Naturally bonded 10Gbps is better.
		* Object storage is more resilient to 1Gbps.
		* Block storage and CephFS are better off on bonded 1Gbps or 10Gbps links.

| Microservice | Resource | Minimum Requirement | Recommended | Notes |
| --- | --- | --- | --- | --- |
| `ceph-osd` | CPU | 1 core | 2 core | 1 core per 200-500MB/s |
| | | | | 1 core per 1K-3K IOPS |
| | | | | ARM may require more cores |
| | RAM | 2-4GB per daemon | 4+GB per daemon | more is better |
| | Disks | 1x per daemon | See notes | Separate partition for OSD from Opsys |
| | DB/WAL | | | |
| | Network |1x1Gb/s | 1x1Gb/s | 10Gb/s bonded for hi-perf |
| `ceph-mon` | CPU | 2 cores | 2 cores | benchmarking needed |
| | RAM | 5GB per daemon | more as scale increases | maps grow with cluster size |
| | Disks| 100GB per daemon | | SSD is recommended |
| | Network |1x1Gb/s | 1x1Gb/s | 10Gb/s bonded for hi-perf |
| `ceph-mds` | CPU | 2 cores | 2 cores | MDS is CPU intensive, single threaded |
| | RAM | 2GB per daemon | 4GB+ per daemon | more RAM needed as cluster grows|
| | Disks| 1GB per daemon | 1GB per daemon | more disk as growth occurs|
| | Network |1x1Gb/s | 1x1Gb/s | 10Gb/s bonded for hi-perf |
| `ceph-mgr` | CPU | 2 cores | 2 cores | benchmarking needed |
| | RAM | 2GB | 2GB | |
| | Disks| ??? | ??? | |
| | Network |1x1Gb/s | 1x1Gb/s |  |




 2. More about those three (3) storage methods...
		1. `Ceph Object Storage`
			1. ...consists of a `Ceph Storage Cluster` and `Ceph Object Gateway` (`RGW`)
				1. The `Ceph Object Gateway`...
					1. ...is an "object storage interface built on top of `librados`."
					2. ...provides a RESTful gateway between applications and Ceph storage clusters.
				2. A Ceph Storage Cluster is a collection of `Ceph Monotors`, `Ceph Managers`, `Ceph Metadata Servers` and `Object Storage Daemons` (OSD).  The cluster receives data from `Ceph Clients`.
		2. `Ceph Block Storage`
			1. ...is provided by python `librbd`
			2. ...works in conjunction with hypervisors (QEMU, Xen, etc.), or an abstraction layer such as `libvirt.`
			3. A `Ceph Block Device` is ...also called "RADOS Block Device" (RBD).  It is "a software instrument that orchestrates the storage block-based data in Ceph."  "Ceph Block Device splits block-based application data into 'Chunks'.  RADOS stores these chunks as objects."
		3. `Ceph File System` (`CephFS`)
			1. A POSIX-compliant file system built on top of the Ceph distributed object store (RADOS).
	7. Ceph is highly scalable and highly available.
		1. This means we can distribute the cluster over many POE-controlled Rock64 nodes.
		2. The storage can be integrated across a heterogenous hardware ecosystem as it scales.
	8. Resource requirements for Ceph are modest (discussed later)
 3. What are the parts of Ceph?
	 1. General architecture 
	 2. Ceph is implemented as a cluster of five(5) microservices:
		 1. `ceph-monitor` (monitor): 
			 1. Maintains maps of the cluster state, including--
				 1. ***monitor map*** -
					 1. contains the cluster `fsid`, the position, the name, the address and the tcp port of each monitor instance.
					 2. specifies the current epoch (time) of the monitor map's creation and last modification.
					 3. (to view a `monitor map` run `ceph mon dump`)
				 2. ***manager map*** -
					 1. contains cluster `fsid`, the position, name, address and tcp port of each manager instance.
					 2. specifies the current epoch (time) of the map's creation and last modification.
				 3. ***OSD map***
					 1. Contains the cluster `fsid`, the time of the OSD map's creation, the time of the OSD map's last modification, a list of pools, a list of replica sizes, a list of PG numbers, and a list of OSDs and their statuses (e.g. `up`,`in`).
					 2. To view an osd map, run `ceph osd dump`.
				 4. ***MDS map***
					 1. Contains the current MDS map epoch when the map was created and the last time it was changed.
					 2. Contains the pool for storing metadata, a list of metadata servers (and their connection info) and their status (`up`,`in`).
					 3. To view an MDS map, run `ceph fs dump`
				 5. ***CRUSH map***
					 1. Contains a list of storage devices and the failure domain hierarchy (for example device, host, rack, row, room) and the rules for traversing the hierarchy when storing data (...or in simpler terms, it contains where data will be stored so it can be stored in diverse places to keep the data safe).
					 2. To view the CRUSH map...
						 1. `ceph osd getcrushmap -o {filename}` for the raw result...
						 2. This raw result is then decompiled using `crushtool -d {comp-crushmap-filename} -o {decomp-crushmap-filename}`
						 3. The decompiled crush map file can then be displayed using `cat`
				 6. ***PG Map (Placement Group)***
					 1. Contains the PG (Placement Group) version, its timestamp
					 2. Contains the last OSD map epoch, the full ratios, the details of each placement group, including--
						 1. ...PG ID
						 2. ...the Up Set
						 3. ...the Acting Set
						 4. ...the state of the PG (example: active + clean)
						 5. ...data usage statistics for each pool
			 2. Manages authentication between clients and daemons.
			 3. <span style="color:red">The cluster will not operate without these maps.</span>(The maps are critical to allow the various Ceph daemons to coordinate with one another. )
			 4. Requires at least three (3) instances for HA.
		 2. `ceph-mgr` (manager):
			 1. Responsible for keeping track of runtime metrics and the current state of the Ceph cluster, including storage utilization, current performance metrics and system load.
			 2. Exposes web-based dashboard and rest-API for exposing cluster information.
			 3. Requires at least two (2) instances for HA.
		 3. `Ceph-osd` (OSD or Object Storage Daemon)
			 1. stores data
			 2. handles replication, recovery and rebalancing of data (see CRUSH)
			 3. provides some monitoring information to ceph-monitor and ceph-mgr nodes
			 4. checks other OSD daemons for a heartbeat.
			 5. Requires at least three (3)OSDs for HA.
			 6. Interacts with logical disks 
	 3. How does Ceph store data?
		 1. Ceph stores data as objects identified uniquely across the cluster balanced across the fleet of available OSDs ![[Pasted image 20231224020311.png]]
		 2. Each OSD handles the storage 
		 3. Ceph OSDs store the data as as a flat system of objects.
		 4. Any hierarchy (such as in a file system) is provided in metadata by the clients.
	 4. What are the points of failure?
		 1. Risk: If the CRUSH / PG / MDS maps become corrupted, ceph may lose TRACK of data.
			 1. Mitigations: 
				 1. Maps are replicated across multiple nodes.
			 2. Unknowns: 
				 1. Are maps ACID?
				 2. Are maps versioned to allow rollbacks?
		 2. Risk: If the client cannot talk to a `ceph-monitor` (if all monitors become unavailable), the clients cannot read from or write to the cluster.
			 1. Mitigation: 
				 1. Multiple monitor instances reduce the likelihood of this.
				 2. Containerized monitor instances should allow faster recovery if this happens.
			 2. Unknowns:
				 1. How long would it take a restarted container instance to come online?
		 3. Risk: What is the authentication mechanism?  Is it weak?
			 1. Mitigation:
				 1. Needs exploration
			 2. Unknowns:
				 1. Needs exploration
	 5. Minimum Viable Product (What is the minimum we need for a local dev environment?)
		 1. Ceph dev environments can run with one monitor but that is a SPOF.

---

---
## Operating System Requirements
1. ubuntu 20.04 and 22.04 are well tested and well supported with Ceph Reef 18.2.xx
2. `btrfs` is not supported on `centos 7`...but who wants `CentOS` given their recent issues with container support?
3. `bluestore` is recommended over `btrfs`.
4. Kernel >5.3 recommended for RBD and CRUSH tunables.
---
## Deployment Options
1. Well supported options (assuming ceph `reef` version):
	1. `cephadm` - a tool for deploying and managing a ceph cluster
		1. Supported only on `Octopus` version or newer
		2. Fully integrated with orchestration API
		3. Requires python3 and docker/podman.
		4. Requires `systemd`
		5. Cephadm works only with BlueStore OSDs.
		6. <span style="color:red">Uses root logins and ssh connections as root user</span>
	2. `rook` - a tool for deploying clusters in Kubernetes
		1. Recommended when deploying in Kubernetes
		2. Recommended to connect existing ceph clusters to Kubernetes
		3. Supported only on `Nautilus` and newer ceph versions
		4. Supports the orchestrator API and management features in the CLI and dashboard.
		5. <span style="color:red">Needs deep-dive security review.</span>
	3. Manual Deployment
		1. Will allow security hardening
		2. Will allow Dockerfile development for our planned use-case.
		3. Avoids the assumptions of `rook` and `cephadm`.
		4. This will require a lot more effort.
2. <span style="color:green"><b>Selected Deployment Option:</b></span>
	1. We will use a manual deployment to aid in the development of a `Dockerfile` with our approved base images to avoid pulling in any questionable practices (like root user ssh connections).
	2. Our planned use-case will deploy container-based ceph nodes to Rock64/Raspberry PI OSD nodes.
	5. We will consider `rook` later for integration of the cluster with Kubernetes when we begin building out ceph services on the Kubernetes cluster for the `datacollectors` project.
---
## Initial Container Development
#### Prerequisites
1. We really need a c++ builder that has boost, etc. already installed to reduce the time required to build ceph.
2. The `monorepo` command needs to build only projects which have changed if a "skip if unchanged" flag is set to true.  This will allow other projects to build/test and ensure forgotten dependencies are tested while allowing the minority of time-consuming projects to skip a rebuild (e.g. building and installing boost) with each change...

#### Building Ceph from sources...


1. Steps that should be in `Manifest.yml`:
	1. Create `ceph` source code submodule: `git submodule add -f --name ceph git@github.com:ceph/ceph.git containers/services/ceph/src/ceph`
```
2. Build this base: 
```bash
docker build --compress --tag ceph:dev -f containers/services/ceph/Dockerfile .
```
4. Run the base image and manually walk through the instructions [ceph docs](https://docs.ceph.com/en/latest/install/clone-source/) and [building-ceph](https://github.com/ceph/ceph#building-ceph)
5. Consider running `cmake -DDIAGNOSTICS_COLOR=always` to keep color coded logging.
6. 