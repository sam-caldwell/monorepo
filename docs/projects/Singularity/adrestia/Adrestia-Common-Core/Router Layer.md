# Objectives
* Provide an HTTP endpoint for health checks between peers.
* Provide an HTTP endpoint for health checks by clients.
* Provide a transparent virtual IP management facility which will run in the background of the application to--
	* discover, authenticate and manage peer connectivity
	* manage the virtual IP address configuration among the peers
* Provide layer 3+ traffic routing from client/peer to the lower layers.

# Startup
1. Node starts and configures ethernet interface with DHCP or other host-specific mechanism with local IP address.
2. Node starts its Router Layer background web service
	1. peer health check endpoints are enabled (`api/v1/router/health/peer`)
	2. client health check endpoints are enabled (`api/v1/router/health/public`)
	3. peer hint endpoint is enabled (`api/v1/router/peering`) (requires authentication by PSK)
3. When a peer hint is given to the node, it will attempt to connect to the peer.
	1. If peer connects, peer will negotiate leader and the cluster map will eventually update.
	2. If peer fails to connect, the hint will be removed.  failed peer hints will be considered for rate limiting of source client.

* When a implementing application starts, it only has its local IP address.
* The application will start its peer health-check endpoints.
	* Will return current node state
* The application will start its client health-check endpoints.
	* Will return http/425 (too early) until peers are online and ready.
* Until at least two peers (total three nodes) are online, the node will remain in a standby though it will continue to discover other peers in the subnet.