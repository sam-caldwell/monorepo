# Tasks
1. Create edge server
	1. Create clusterMap configuration file
		1. file validator
		2. file loader
	2. Extend clusterMap with runtime state
	3. Create HA network interface functionality
		1. Implement VIP failover functionality
		2. Implement heartbeat functionality
		3. extend heartbeat with stats in clusterMap.
	4. Create test backend server to respond to edge proxy requests
	5. Create test client  to send requests to edge.
2. Create Test Client
	1. Create object (HTTP/PUT)
	2. Read object (HTTP/GET)
	3. Update object (HTTP/POST)
	4. Delete object (HTTP/DELETE)
	5. List objects (HTTP/LIST) (Custom method)
2. Create Adrestia Daemon
	1. HTTP server
	2. Mock responder
6. Create Meta Server
	1. HTTP server
	2. Mock responder
	3. Implement VIP
7. Create Data Object Server
	1. Adrestia Daemon object cache (noop)
	2. Adrestia Daemon metadata cache (noop)
	3. Adresita Daemon forward request to edge.