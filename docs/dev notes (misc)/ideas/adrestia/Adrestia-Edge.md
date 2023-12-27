# Objectives
* Provide a service which will sit at the perimeter of an "Availability Zone" and receive any requests from the local daemon instances used by clients to create, fetch, update or delete objects.
* Provide a service which will expose metrics and state information about the health of the Availability Zone (known as the "clusterMap"), enabling client instances and other edge nodes to verify the health of peer AZs.
* Provide a service which will allow edge nodes within an Availability Zone (AZ) to replicate (push) clusterMap changes to its peers.
* Provide a service which will authenticate connections between peer edge nodes and between clients and availability zones.
* Provide a service which will allow network failover of a virtual IP address.

# Authentication
* Any node (edge or otherwise) will authenticate with other nodes in the cluster using TLS client and server-side certificates issued by the PKI service.  Each Availability Zone (AZ) must have a PKI service instance.
