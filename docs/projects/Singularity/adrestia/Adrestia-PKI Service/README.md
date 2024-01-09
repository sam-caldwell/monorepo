# Overview
 The public-key infrastructure (PKI) service exists as a cluster of one or more PKI Service instances per availability zone (AZ) to create and maintain a chain of Certificate Authority (CA), Intermediate Certificate Authority (ICA) and Certificate Revocation List (CRL) objects for the Adrestia cluster.
# Definitions

| Term | Definition |
| ---- | ---- |
| PKI Service | The public-key infrastructure service used to maintain a certificate authority (CA), intermediate certificate authority (ICA) and certificate revocation list (CRL) for the Adrestia ecosystem. |
| Availability Zone | An autonomous storage zone (rack, building, data center) with some degree of isolation from the rest of the ecosystem. |
| PKI Cluster | A federation of PKI service instances across multiple Availability Zones (AZs) sharing a common clusterId (UUID) and virtual IP address (`pkiGlobalVip`), represented by an elected `pkiGlobalLeader` node. |
| Local PKI | The PKI service instances sharing a common Virtual IP address (`pkiLocalVip`) and common Availability Zone (AZ), represented by an elected `pkiLocalLeader` node. |
| `pkiGlobalVip` | A virtual IP address representing the PKI cluster spanning one or more AZs. |
| `pkiLocalVip` | A virtual IP address representing the PKI cluster within a single AZ |
| `pkiGlobalLeader` | The peer-elected cluster-wide leader responsible for maintaining the PKI cluster map and `pkiGlobalVip`. |
| `pkiLocalLeader` | The peer-elected AZ-specific leader responsible for maintaining the AZ-level copy of the PKI cluster map and `pkiLocalVip`. |
| Cluster Map | A data structure representing the current configuration of the PKI cluster, which is maintained by `pkiGlobalLeader` and `pkiLocalLeader` and replicated to each PKI instance through the health check process. |
|  |  |

# Requirements ![[cluster-topology-drawing]]
1. **Minimum Viable PKI Cluster**
	1. For each Availability Zone (AZ) there must be at least one (1) PKI service instance online to serve the needs of the Adrestia ecosystem.
	2. The PKI cluster must have at least one (1) `pkiLocalLeader` per AZ.
	3. The PKI cluster must have at least one (1) `pkiGlobalLeader` for the ecosystem.
3. **PKI Node Initialization** 
	1. **Phase I: Leader Discovery**
		1.  **Goals**:
			1. Confirm the availability of `pkiGlobalLeader` or self-elect if not found after `pkiInitRetry` retries.  If `pkiGlobalLeader` not found then the new node becomes both `pkiGlobalLeader` and `pkiLocalLeader`.
			2. Confirm the availability of `pkiLocalLeader` or self-elect if not found after `pkiInitRetry` retries. This means `pkiGlobalLeader` is known but `pkiLocalLeader` is not known and the new node assumes that role.
		2. **Caveats**
			1. The requests in this stage should all be insecure HTTP requests.
			2. This means that any endpoint could discover the `pkiGlobalLeader` or `pkiLocalLeader` and their virtual IP addresses.
		3. **Drawing** ![[phase1-initialization-drawing]]
	2. **Phase II: Node Registration** 
		1.  **Goals**
			1. Get the root and intermediate Certificate Authority certificates (CA, ICA) for any existing `pkiGlobalLeader` and/or `pkiLocalLeader`.
			2. This stage will authenticate the new node using a pre-shared key (PSK) and start the authentication dance for future requests between nodes.
			3. This stage will allow the node to get the root CA certificate and intermediate CA (ICA) certificate for the cluster and AZ, allowing HTTP requests thereafter.
			4. This will permit the new node to sign new certificates.
		2. **Caveats**
			1. The first request in this stage is insecure HTTP.
			2. The second request is encrypted using HTTPS and authenticated.
		3. **Drawing **![[phase2-initialization-drawing]]

	
	3. Node Startup
		1. The node starts with the following:
			1. A pre-shared key (`authPsk`) used to authenticate requests.
			2. A `pkiClusterVip` (virtual ip address).
		2. The node sends  `HTTP/GET` to `http://${pkiClusterVip}/api/v1/status`.
			1. The node times-out because no cluster leader responds
			2. If the request times-out, it is retried `pkiInitRequestRetry` times. 
				1. If the request time-out exceeds retries, the node will know there is no `pkiGlobalLeader` and proceed to "Cluster Initialization".
			3. If there is an active `pkiGlobalLeader` it will--
				1. Use `REMOTE_ADDR` (client IP address) to identify the node's availability zone (AZ).
				2. redirect the request to `pkiLocalVip` for the node's Availability Zone (AZ) based on the `REMOTE_ADDR` (client IP).
			4. 
		4. If the node receives `http/200 OK` in response, the node may assume the request is redirected to `http://${pkiLocalVip}:80/api/v1/status`.
		5. If the node receives `http/404` in response, the AZ has no `pkiLocalLeader` and the node will immediately assume `pkiLocalLeader` for the given AZ (See "Local Leader Initialization").
		6. it will proceed to "Member Initialization," below.
		7. If the request times-out, it will be retried `pkiInitRequestRetry` times.
		8. If the request retry exceeds `pkiInitRequestRetry`, the node will proceed to "Cluster Initialization."
	4. "Cluster Initialization"
		1. The node sends `http/get` request to `http://${pkiClusterVip}/api/v1/ca`
		2. If the request to `http://${pkiClusterVip}:80/api/v1/status` returns `http/200 OK`, 
		3. Assume it is the `pkiGlobalLeader` and `pkiLocalLeader`.
		4. Takeover the `pkiGlobalVip`
		5. Takeover the `pkiLocalVip`
		6. Start serving the `api/v1/clusterMap` endpoint over insecure HTTP.
7. The PKI service will operate as a "Global" cluster across all Availability Zones (AZ) in an Adrestia deployment.  The global cluster requires that--
	1. Each local "availability zone" (AZ) will have at least one (1) online public-key infrastructure (PKI) service instance running.
	2. Each local AZ must have a single virtual IP address (VIP) known as the PKI_LOCAL_IP.
	3. Each AZ will share a PKI_GLOBAL_IP virtual IP address (VIP) representing the entire cluster.
	4. 
	
	
2. 
3. 
4. For each Availability Zone (AZ) there must exist at least one (1) public-key infrastructure (PKI) service instance.
5. 
6. Each PKI server instance will act as a node in a global PKI cluster spanning all Availability Zones.
7. Each PKI server will share an AZ-wide IP address for communications within the Availability Zone (AZ).
8. The PKI cluster will--
	1. Be identified by UUID (`clusterId`).
	2.  Represent all PKI nodes in aggregate across the system with an elected leader
		1. Cluster (Global) Leader
			1. Only one for the cluster.
			2. Maintains the root CA.
			3. Performs health checks on AZ leader nodes.
			4. Holds the global virtual IP address (VIP) for the cluster.
			5. Presents the full API to any client (may shed load to local leader VIPs).
			6. Maintains the authoritative copy of the cluster map.
			7. Maintains the authoritative copy of the certificate revocation list (CRL).
		2. AZ (Local) Leader
			1. Has its own ICA and the CA cert.
			2. May be elected Global Leader if needed.
			3. Each AZ has one.
			4. Maintains a node ICA (like any node)
			5. Performs local health checks on AZ nodes.
			6. Responds to global leader health checks to update global cluster map.
			7. Holds the local virtual IP address (VIP) for each Availability Zone (AZ).
			8. Presents the full API to any client (may shed load to member nodes).
			9. Maintains the authoritative copy of the local cluster map and revocation list (CRL), sync'd with the Global leader.
9. Clients may request services from either a Global VIP or local (AZ) VIP but they should send requests to the local AZ VIP where appropriate.
10. Cluster and AZ leaders may redirect requests or otherwise shed load to an appropriate AZ node.  For example--
	1. Assume there is an AZ on 10.1.1.0/24 and 10.2.1.0/24.
	2. Assume there is a PKI global leader on 10.2.1.10 and a PKI local leader on 10.1.1.10.
	3. A request from 10.1.1.100 to 10.2.1.10 (global leader) should be redirected to the local leader 10.1.1.10 rather than answered directly.
	4. If a new AZ at 10.3.1.0/24 is added (assuming local leader at 10.3.1.10, requests from this subnet would likewise be shed to the local leader on that subnet.
	5. This means that if an AZ local leader is not available, the request must be retried until one becomes available.  But the cluster leader or other AZ local leaders should not begin responding to requests for other AZs.
	6. This also means that if 10.3.1.100 sends a request to 10.1.1.10 (a local leader on a different AZ), the request should be redirected to 10.30.1.10 (the local leader on the same AZ).
11. 