* **General**
	* The client daemon, or "Adrestia-Daemon" runs on a local client host and provides connectivity, caching, authentication and authorization between the client and the [[Adrestia-Cluster]].
	* The daemon is intended to run in the background on the client host at all times in order to maintain its awareness of the cluster's VIP, AZ VIPs, performance tuning, and both object and metadata caches.
* **Authentication and Authorization**
	* The [[Adrestia-Daemon]] provides client-to-server authentication and authorization using certificates issued by the cluster public-key infrastructure (PKI) service.
	* A client is presumed to be authorized to access any object it has created or for which it possesses knowledge.  This assumption is made because all objects are encrypted, and the encryption keys must be shared for them to be readable by the client application.
	* The [[Adrestia-Daemon]] receives its authentication/authorization PKI certificates by submitting an "Adoption Request" to [[docs/dev notes (misc)/ideas/adrestia/Adrestia-PKI Service/Requirements]].  
		* This "Adoption Request" comes in the form of a certificate-signing request (CSR).
		* When the request is approved, the PKI service will sign the CSR and return the signed certificate.
		* The [[Adrestia-Daemon]] will use the signed certificate for all subsequent communications with the cluster.
* **Object Signing and Encryption**
	* The [[Adrestia-Daemon]] is responsible for signing and encrypting any objects it may create using client-side encryption keys.
	* The [[Adrestia-Daemon]] may publish a public encryption key to the PKI service to facilitate object sharing between clients, where a remote client will encrypt/sign an object using another client's published cryptographic key to allow the foreign client to access the shared object.
	* An object stored in the [[Adrestia-Cluster]] is thus nothing more or less than a blob of encrypted bytes once received by the cluster [[Adrestia-Edge]].
	* Once signed and encrypted, the object should only be identifiable by its [[Metadata]] and[[ObjectId]].
* **Object and Metadata Caches**
	* Any object or metadata obtained from the cluster using an instance of [[Adrestia-Daemon]]]] must be presumed to be compromised because the daemon has the ability to encrypt/decrypt/sign objects under its jurisdiction.
	* Encryption/Decryption and signing are expensive operations.
	* To offset the expense of encryption/decryption, the [[Adrestia-Daemon]] maintains two caches:
		* ***OBJECT CACHE***: An encrypted local cache of frequently-used objects with a configurable expiration window.
		* ***METADATA CACHE:*** An encrypted local cache of frequently-used meta data records with a configurable expiration window.
	* **Configurable Cache Expiration**
		* Data published into the caches using a rotating set of symmetric encryption keys.  As keys are expired off the key table, cached objects stored using the expired keys are rendered unreadable and marked for garbage collection.
		* The rate at which symmetric cache encryption keys are expired is configured first by cluster policy but may be adjusted to a shorter time period by the [[Adrestia-Daemon]]]].  *Note: The cluster must always presume the longer expiration window and not trust any assertion of a shorter window by a client, as this could be deceptive.*
	* Cache functionality should be defined in [[Encrypted-Cache]], a reusable Golang library in the monorepo.
* **ClusterMap**
	* The client [[Adrestia-Daemon]] will maintain a sanitized version of the [[Adrestia-ClusterMap]] which is obtained from the [[Adrestia-Edge]] during initialization and updated periodically as part of the health-check requests.
	* The `ClusterMap` will identify only the following information to the [[Adrestia-Daemon]]:
		* `ClusterId` (UUID)
		* Connection information
			* Cluster Virtual IP address and port (ClusterVIP).
			* Each Availability Zone (AZ) Virtual IP address and port.
		* Cluster state
			* Overall cluster state (up, down, degraded).
			* Per AZ cluster state (up, down, degraded).
		* Performance and Capacity Metrics
			* All metrics should be subdivided into 300s, 1800s, 3600s, lifetime windows.
			* All metrics should be aggregated to the cluster level.
			* All metrics should also be broken down by availability zone.
			* Performance:
				* min, max, average and p99 latency on object and metadata read/write/list operations.
				* hit and error counts
				* rebalancing operation count and average duration
			* Capacity:
				* Number of Objects Stored
				* Minimum, Maximum and Average Object Size
				* Number of Metadata records
				* Minimum, Maximum and Average Metadata Key size
				* Minimum, Maximum and Average Metadata Value size
	* The [[Adrestia-Daemon]] should conserve network bandwidth when performing health check requests to obtain the latest `ClusterMap`.
		* When the daemon starts, it may request the full `ClusterMap` from the edge.
		* After starting, the daemon should only request the latest changes to the `ClusterMap` by providing the `lastUpdated` timestamp (uint64) to the cluster.
		* The cluster should reply to a `lastUpdated` timestamp with a document containing only the records that changed between the current `ClusterMap` and the when the daemon's version was last updated.

