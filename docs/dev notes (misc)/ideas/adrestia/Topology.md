![[Topology-Drawing]]

# Components
![[Client Applications]]
![[Adrestia-Daemon]]
![[Adrestia-Availability-Zone]]


# Edge Nodes
1. An "edge node" represents the perimeter of an AZ using VRRP to present a single virtual IP address (VIP) to which clients connect.
2. The edge node routes client requests to the appropriate object or metadata storage node.
3. The edge node is identified by a 64-bit `zoneId`.
# Meta Nodes
1. A cluster must have one or more Meta nodes per Availability Zones (AZ). ![[Pasted image 20231226123609.png]]
2. Each Meta Node is identified by UUID (`NodeId`).
3.  Each "Meta Node" stores zero or more key-value records about a given object, where the record is identified by `recordId`:

		recordId:=hash(ObjectId + KeyString)

4. The `recordId` (hash) is used to balance the metadata records across the meta nodes.
	1. Given `n` nodes, the key-space is divided across `n` nodes evenly.  Each node has in its clusterMap a starting and ending hash.
	2. As more nodes are added to the cluster, this geometry can be changed by updating the starting and ending hash for a given node.  Any records outside this boundary must then be rebalanced to other nodes.
	3. Given the clusterMap's knowledge of which Meta node has which range of hash values, the edge nodes can quickly route a request to the appropriate node.
5. Meta Node key-value records must meet the following expectations:
	1. The Key may be up to 1,024 bytes in length
	2. The Value may be up to 65,536 bytes in length.
	3. The Key is prefixed by the ObjectId of the related object.
	4. The key-value record is identified by the hash of the ObjectId and key:
# Data Nodes
1. A cluster must have one or more Data nodes per Availability Zone (AZ). ![[Pasted image 20231226131113.png]]
2. Each Data Nodes is identified by UUID (`NodeId`).
3. Each "data node" consists of two (2) file partitions--
	1. ***index partition*** - stores a table of ObjectId (hash), ObjectSize (uint64) and position (uint64) as a binary tree.  This tree allows an object to be found quickly.
	2. ***blob partition*** - stores blobs of object data described by the index partition.
4. Each record in the `index partition` has a `lock` flag, allowing the record to be locked during replication or rebalancing operations.
5. A Data Node must have a counterpart in at least two availability zones for HA.
6. When data is written to a single data node on one AZ, the data is edge node will not confirm the operation until a quorum of data nodes agree that the write is complete.
7. When in steady state (i.e. all data is replicated), the bit-level image of each Data Node should be identical on both the index and blob partitions.
# ObjectId
  1. An ObjectId is a SHA-512 hash, creating a block chain, deriving from a common seed.
  2. The "common seed" is derived from the ClusterId (UUID):

	        commonSeed:=HMAC( clusterId, clusterPrivateKey )
            objectId:=HMAC( concat( seed, objectId), clusterPrivateKey)

  3. All objects should derive their ObjectId from this common seed EXCEPT where an object is a new version or other logical derivative of some other object.  In these cases, the "seed" is the ObjectId of the object which is related. This creates a mathematically provable chain.
  4. The purpose of the ObjectId is to--
	   a. Uniquely identify all objects across the cluster.
	   b. Ensure that object identification within one cluster can be mathematically sequestered from other clusters.
# NETWORK CONNECTIVITY
1. Each client must be able to connect to the edge nodes of at least one availability zone.
2. While a client may not connect to all of the availability zones, this could diminish the client's ability to survive an AZ failure.
3. All Availability Zones must be able to connect to all other Availability Zones since each Edge, Meta and Data Node must be able to replicate data to their respective peers.
  4. When a node comes online...
	  1. Its IP address should be issued by DHCP.
	  2. The device role will be determined by clusterMap, which will map host MAC address to cluster role.
# CLUSTER MAPS
1. Each node is provisioned with a ClusterMap.yml file which identifies--
	1. The hostname, MAC address, role, IP address and port numbers.
	2. A mapping of each node in the cluster to it's respective availability zone.
	3. The public and private encryption key files used by the cluster for in-flight and at-rest encryption.
	4. The ClusterID (uuid).
	5. The readQuorum number (the number of peers which must agree before an answer is returned to the client)
	6. The writeQuorum number (the number of peers which must confirm an object or metadata write before it is acknowledged).
	7. The azQuorum (the minimum number of Availability Zones which must be online for the 
2. Once a host is online, the ClusterMap will persist in memory along with additional runtime state (metrics)
	1. The runtime-state (metrics) maintained by a host will be used to route messages effectively to reliable peers first and less reputable peers thereafter.
	2. Each host will only maintain the clusterMap state needed to perform its specific role.
	3. The runtime state of a host will include--
		1. host state (up/down)
		2. performance metrics from heartbeat
			1. firstSeen, lastSeen timestamps
			2. 300s, 1800s, 3600s moving average response time of each host.
			3. 300s, 1800s, 3600s error count of each host.
			4. 300s, 1800s, 3600s ok count of each host.
		3. performance metrics from host-specific payload actions
			1. firstSeen, lastSeen timestamps
			2. 300s, 1800s, 3600s moving average response time of each host.
			3. 300s, 1800s, 3600s error count of each host.
			4. 300s, 1800s, 3600s ok count of each host.
	4. Each host will use its runtime metrics to determine the order of preference when deciding how to replicate information. 
# Public Key Infrastructure
1. Each Availability Zone (AZ) must have a public key infrastructure (PKI) service instance.
2. The PKI instance will--
	1. Issue a Certificate Authority (CA) for the cluster.
	2. Issue an Intermediate CA (ICA) for each Availability Zone.
	3. Issue certificates signed by the ICA of each Availability Zone (AZ) for each node in the cluster.
	4. Issue Client Certificates to any Adrestia Daemon instance registering itself with an Edge node.
	5. Publish/expose a certificate revocation list (CRL) used to evict nodes/clients from the cluster.
	6. Expose an API for requesting, signing and revoking PKI certificates.
3. During cluster provisioning, the PKI instance for an AZ must be the first instance deployed, and during deployment this PKI instance will establish the "trust root" (CA) and ICA for the cluster.
