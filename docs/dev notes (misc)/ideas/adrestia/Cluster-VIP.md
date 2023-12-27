Each cluster has a single IP address which is passed between edge nodes elected as "cluster master."  This IP address is the endpoint to which all client requests are sent.

* A cluster master is elected when--
	* The first edge node comes online and assumes the Cluster Master role.
	* Two or more edge nodes elect one node as the Cluster Master.
* Re-election of cluster master may occur when--
	* The current cluster master decides to step down (e.g. when shutting down or shedding load).
	* A quorum of edge nodes determine that a new cluster election is needed because the current master is inaccessible or unhealthy.

