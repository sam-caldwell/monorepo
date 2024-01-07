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
		* [[Adrestia-Data]] nodes which provide object storage, replication and re-balancing services for the Availability Zone.
		* [[docs/dev notes (misc)/ideas/adrestia/Adrestia-PKI Service/Requirements]] nodes which provide public-key cryptography services for the Availability Zone and cluster.
* **Capacity versus Redundancy**
	* The Availability Zone provides no data redundancy
	* Each [[Adrestia-Meta]] and [[Adrestia-Data]] node stores a unique set of data entities.
	* Adding [[Adrestia-Meta]] and [[Adrestia-Data]] nodes to an Availability Zone (AZ) only increases the capacity to store more data and to distribute load for creating new data records or accessing existing records across a wider fleet of autonomous systems.
		* Fewer nodes with larger disks is discouraged.
			* This has a lower cost per GB-hour but a lower performance expectations as there are fewer nodes to process data operations.
			* This also creates a higher risk of a single failure affecting more data and thus requiring longer and more CPU/network intensive re-balancing operations.
		* Many smaller, lower costs nodes is recommended.
			* While this increases the cost per GB-hour due to additional power required and the cost per GB of smaller versus larger disks, the impact of a single node failure will be smaller in terms of affected data and recovery/re-balancing efforts.
			* A larger fleet of nodes will defray the cost of concurrent operations across many independent nodes, increasing performance.
* Fault Tolerance and High Availability (Inter-AZ Replication)
	* Edges and Virtual IP Addresses
		* The cluster is made fault tolerant when there exist more than one (1) Availability Zone (AZ) which is isolated physically and in terms of power and network access by clients from the other Availability Zones (AZs).
		* It is recommended to run at least three (3) Availability Zones per cluster with additional AZs providing further data redundancy at diminished marginal costs.
		* HA and Fault Tolerance are facilitated through the [[Cluster-VIP]] which can pass between Availability Zones as edge nodes deem appropriate to ensure availability.  
		* The edge node with the [[Cluster-VIP]] is the elected cluster leader.
		* Within the Availability Zone (AZ) may exist one or more [[Adrestia-Edge]] Nodes which share an [[AZ-VIP]] which can be passed between the edge nodes of a given AZ as they deem appropriate to ensure AZ availability.  
		* The edge node with the [[AZ-VIP]] is the elected AZ leader.
		* Each Cluster and AZ leader will poll the health check endpoints of its member edge nodes and update the clusterMap with their reported statuses.
			* Each member node will observe the health check requests as evidence that the leader is alive.
			* Allowing the leader to poll members reduces the network traffic required to health check the cluster.
	* Data Redundancy and Replication
		* There are three (3) node types that require data replication:
			* [[docs/dev notes (misc)/ideas/adrestia/Adrestia-PKI Service/Requirements]] nodes - store/maintain encryption keys.
			* [[Adrestia-Meta]] nodes - store/maintain key-value metadata records.
			* [[Adrestia-Data]] nodes - store/maintain user data objects.
		* Each node type has different minimum requirements:
			* [[docs/dev notes (misc)/ideas/adrestia/Adrestia-PKI Service/Requirements]] 
				* Must have at least one (1) instance per AZ.
				* If PKI instance is not available, the cluster cannot create/sign new certificates or revoke existing certificates.
				* Data objects and metadata are still available for read/write operations.
			* [[Adrestia-Meta]]
				* Must have at least one (1) instance per AZ.
				* If metadata service is not available, object search by metadata is no longer possible.
				* Data objects are still available for read/write assuming the `ObjectId` is still known.
				* Adding more metadata service nodes per AZ will increase the total storage capacity and load capacity of the cluster.
				* Metadata service nodes must be added to the cluster in cryptographically paired nodes across the AZs.  For example, given three AZs, the metadata service must be scaled by adding three nodes at a time to maintain a balanced replication.
			* [[Adrestia-Data]]
				* Must have at least one (1) instance per AZ.
				* If data service is not available, data objects cannot be read/written.
				* Adding more data object service nodes per AZ will increase the total storage capacity and load capacity of the cluster.
				* Data object service nodes must be added to the cluster in cryptographically paired nodes across the AZs.  For example, given three AZs, the data object service must be scaled three nodes at a time to maintain a balanced replication.
		* **Joining the Cluster:** When a replicating node is added to the cluster--
			* The node must be adopted by the [[docs/dev notes (misc)/ideas/adrestia/Adrestia-PKI Service/Requirements]] instance.
				* This will result in the node receiving its signed certificate pairs and the cluster CA/ICA certificates.
				* At this point, the node will be added to the ClusterMap.
				* The edge nodes will learn about the node from this ClusterMap registration.
			* Next, the node must establish itself within the pool of node types.
				* This is an assertion in the node's Certificate and should appear in the ClusterMap.
				* If the edge nodes have no other registered nodes of this node type, the new node will become the initial node for the given node type.
				* if there is only one AZ in the cluster, no further action is needed.
			* If the node is joining as a member of an existing node type pool, the node must assert the identity of it's replication peer(s).
				* A node can only replicate to a peer in another Availability Zone.
				* A node can only have one peer in each Availability Zone.
			* Once a node is joined to the Availability Zone and it has a joined peer in each Availability Zone to meet the quorum count, the node will be marked as "online" and "ready."
				* Given a quorum count of (q) and an AZ count of n, n must be greater than q and the number of nodes for a given meta/data node type must be greater than q.
		* **Health check and master-member roles**
			* Each member of a node-type peer group exists in its own AZ.
			* Each member of the peer group will hold one of two roles:
				* **Master**
					* Responsible for performing health checks.
					* Responsible for maintaining and serving the node-type clusterMap.
				* **Member**
					* Observes health check requests from master as evidence the master node is online and healthy.
					* May call for a new master election by sending requests to peers.
					* Performs leader election when a quorum of calls for an election are met.
		* **Data Replication between peer groups**
			* Each member of a node type pool can receive an operation request from the edge nodes in its AZ.
				* **Read operations:** Where this operation reads information from the node--
					* If the relevant data record is not marked as "locked," the record will be returned.
					* If the request has a timeout, the system will wait the timeout period and return the relevant record once the lock is removed.
					* if the request exceeds the system-defined timeout, an error will be returned.
				* **Write operations:** Where this operation writes information to the node store, the record will be written locally to the node's storage device, marked as locked and transmitted to the members of the node's peer group.
					* Once a quorum of nodes in the peer group have acknowledged the write, the originating node will respond to the initial request that the write was successful.
					* Once all online nodes in the peer group acknowledge the write operation, the data record will be unlocked.
					* If the write operation exceeds the system-defined timeout before a quorum of peers acknowledge the write, the write operation will fail and the client will be notified (via edge proxy).
		* **Node Initialization:**
			* When a replicating node comes online within a peer group--
				* its state is Online-Standby
				* The node will select a peer in the "online-ready" state  and request it's "inventory map" (a paginated map of its records).
				* The initializing peer will iterate over the map and store each record in the map to local storage.
			* During initialization, if the initializing node receives any replication messages from a peer due to ongoing operations, the replication message will be appended to the inventory map and processed at the end of the initialization process.
			* Once initialized, the node will remain in "online-standby" for a period of time to "settle" before transitioning to "online-ready."
	* **Node states**
		* Each node one of the following states:
			* **Online-Ready**: 
				* The node has connectivity.
				* The node is ready for operations.
			* **Online-Standby**: 
				* The node has connectivity.
				* The node is not ready for use.
			* **Online-Degraded**: 
				* The node has connectivity.
				* The node should not be sent new operations until it recovers.
			* **Online-Shedding**:
				* The node has connectivity.
				* The node should not be sent new operations.
				* The node is currently shedding any pending jobs.
			* **Offline**: 
				* The node does not have connectivity.
			* **Evicted**: 
				* The node has been evicted from the cluster.
				* The node will not be brought back online.
				* Any cryptographic keys are invalidated.
	* 