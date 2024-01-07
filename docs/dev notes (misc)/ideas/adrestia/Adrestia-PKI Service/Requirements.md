1. 1. For each Availability Zone (AZ) there must exist at least one (1) public-key infrastructure (PKI) service instance.
2. Each PKI server instance will act as a node in a global PKI cluster spanning all Availability Zones.
3. Each PKI server will share an AZ-wide IP address for communications within the Availability Zone (AZ).
4. The PKI cluster will--
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
5. Clients may request services from either a Global VIP or local (AZ) VIP but they should send requests to the local AZ VIP where appropriate.
6. Cluster and AZ leaders may redirect requests or otherwise shed load to an appropriate AZ node.  For example--
	1. Assume there is an AZ on 10.1.1.0/24 and 10.2.1.0/24.
	2. Assume there is a PKI global leader on 10.2.1.10 and a PKI local leader on 10.1.1.10.
	3. A request from 10.1.1.100 to 10.2.1.10 (global leader) should be redirected to the local leader 10.1.1.10 rather than answered directly.
	4. If a new AZ at 10.3.1.0/24 is added (assuming local leader at 10.3.1.10, requests from this subnet would likewise be shed to the local leader on that subnet.
	5. This means that if an AZ local leader is not available, the request must be retried until one becomes available.  But the cluster leader or other AZ local leaders should not begin responding to requests for other AZs.
	6. This also means that if 10.3.1.100 sends a request to 10.1.1.10 (a local leader on a different AZ), the request should be redirected to 10.30.1.10 (the local leader on the same AZ).
7. 