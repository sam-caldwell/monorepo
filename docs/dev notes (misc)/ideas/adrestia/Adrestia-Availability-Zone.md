* **General** ![[Adrestia-Availability-Zone-drawing]]
	* The Adrestia cluster consists of one (1) or more autonomous "Availability Zones" (AZs).
		* The Availability Zone is a logical entity represented by-
			* a `zoneId` (UUID), 
			* an [[AZ-VIP]] (virtual IP address) and
			* a [[Cluster-VIP]] (virtual IP address)  if elected leader by the other AZs.
		* Each AZ is presented to clients through one (1) or more [[Adrestia-Edge]] nodes.
		* Each edge node presents an API to clients through which data access and management operations is facilitated.
	* The Availability Zone represents a single unit non-redundant of object and metadata storage, much like a single disk in a RAID.
	* Multiple AZs provide data redundancy and fault tolerance.
	* Each Availability Zone (AZ) has six (6) parts:
		* [[Cluster-VIP]] (virtual IP address) - An IP address shared across all AZs for a given cluster.
		* [[AZ-VIP]] (virtual IP address) - An IP address shared by all "edge nodes" in the availability zone.
		* [[Adrestia-Edge]] nodes which provide front-end authentication, routing, and exposed API endpoints for the Availability Zone.
		* [[Adrestia-Meta]] nodes which provide metadata search and storage functionality for the Availability Zone.
		* [[Adrestia-Data]] nodes which provide object storage, replication and rebalancing services for the Availability Zone.
		* [[Adrestia-PKI]] nodes which provide public-key cryptography services for the Availability Zone and cluster.
* **Capacity versus Redundancy**
	* The Availability Zone provides no data redundancy
	* Each [[Adrestia-Meta]] and [[Adrestia-Data]] node stores a unique set of data entities.
	* Adding [[Adrestia-Meta]] and [[Adrestia-Data]] nodes to an Availability Zone (AZ) only increases the capacity to store more data and to distribute load for creating new data records or accessing existing records across a wider fleet of autonomous systems.
		* Fewer nodes with larger disks is discouraged.
			* This has a lower cost per GB-hour but a lower performance expectations as there are fewer nodes to process data operations.
			* This also creates a higher risk of a single failure affecting more data and thus requiring longer and more CPU/network intensive rebalancing operations.
		* Many smaller, lower costs nodes is recommended.
			* While this increases the cost per GB-hour due to additional power required and the cost per GB of smaller versus larger disks, the impact of a single node failure will be smaller in terms of affected data and recovery/rebalancing efforts.
			* A larger fleet of nodes will defray the cost of concurrent operations across many independent nodes, increasing performance.
* Fault Tolerance and High Availability (Inter-AZ Replication)
	* Edges and Virtual IP Addresses
		* The cluster is made fault tolerant when there exist more than one Availability Zone (AZ) which is isolated physically and in terms of power and network access by clients from the other Availability Zones (AZs).
		* It is recommended to run at least three (3) Availability Zones per cluster with additional AZs providing further data redundancy at diminished marginal costs.
		* HA and Fault Tolerance are facilitated through the [[Cluster-VIP]] which can pass between Availability Zones as edge nodes deem appropriate to ensure availability.  
		* The edge node with the [[Cluster-VIP]] is the elected cluster leader.
		* Within the Availability Zone (AZ) may exist one or more [[Adrestia-Edge]] Nodes which share an [[AZ-VIP]] which can be passed between the edge nodes of a given AZ as they deem appropriate to ensure AZ availability.  
		* The edge node with the [[AZ-VIP]] is the elected AZ leader.
		* Each Cluster and AZ leader will poll the health check endpoints of its member edge nodes and update the clusterMap with their reported statuses.
			* Each member node will observe the health check requests as evidence that the leader is alive.
			* Allowing the leader to poll members reduces the network traffic required to health check the cluster.
* 
		* [[Intra-Az-Replication]] is facilitated by the [[Adrestia-Meta]] and [[Adrestia-Data]] nodes themselves with their cryptographically paired nodes in other Availability Zones.
			* For a single-AZ cluster, there should be one or more [[Adrestia-Meta]] nodes and one or more [[Adrestia-Data]] nodes.
			* If a second or subsequent AZ is added to the cluster, the AZ must consist of an identical number of [[Adrestia-Meta]] or [[Adrestia-Data]] nodes and each node must be paired with a corresponding node in the existing AZ.
			* The number of [[Adrestia-Meta]] or [[Adrestia-Data]] nodes (and thus AZs) in a given pairing must always be greater than or equal to the cluster's quorum requirements.
				* If the number of nodes is less than quorum, the cluster will no longer be able of reading or writing information.
				* If the number of nodes exceeds the quorum count, the redundancy of data replicas will increase cluster cost, but should improve performance as load can be balanced across more devices.
				* If the number of nodes exceeds the quorum count, the cluster will be capable of withstanding failures.
				* Each cluster should, therefore have at least quorum + 1 availability zones.
				
	* The Availability Zone may be in one of several states--
		* **leader** - The leader has the AZ VIP and Cluster VIP (shared among all AZs)
		* **member** - The Availability Zone is inaccessible or marked as unavailable and cannot be relied upon.				
		* **unavailable** - 	The Availability Zone is inaccessible or marked as unavailable and cannot be relied upon.
		
	* Each Availability Zone (AZ) shares a common [[Cluster-VIP]] which will be serviced by an elected leader among the pool of AZs, and which will failover as needed when--
		* An edge node becomes unavailable or degraded
		* An edge node determines it must shed load for the health of its AZ or the cluster overall.
	* Each 



1. An "Availability Zone" (AZ) is a collection of edge, metadata and data nodes representing a single rack, room, building or other physical location.
2. A cluster may have one or more Availability Zones (AZs), where each AZ has its own set of edge nodes as well as  backing data and meta nodes.         ![[Pasted image 20231226123506.png]]
3. Data within a meta and data nodes are unique to the availability zone.
4. Data is replicated between corresponding meta and data nodes across AZs.
5. Edge nodes share a virtual IP address (VIP) for each availability zone.  But each AZ VIP is unique to the specific AZ.
6. Edge nodes replicate cluster map information across availability zones.