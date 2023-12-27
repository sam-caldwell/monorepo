* General
	* Mission-specific tools which interact with the [[Adrestia-Daemon]] through the [[Adrestia-Common]] library to create, fetch or delete objects in the system.
	* The client application ***REQUIRES*** the [[Adrestia-Daemon]] to be running in order to complete its requests.  If the daemon is not running--
		* The client application will attempt to start the daemon.
		* The client will return an error if it is unable to start the daemon.
		* The client will respond slower if it starts the daemon as there will be no cache and the daemon will first need to be initialized.
	* Client applications only know that objects ***MAY*** exist, but they do not know where or if they actually exist until the [[Adrestia-Daemon]] responds.
* Example Client Applications
	* [[Adrestia-s3]]
		* The `adrestia-s3` application is planned as an s3-compatible frontend for the project which will react to S3 requests and process them through the [[Adrestia-Daemon]] transparently.
		* This client will operate as a micro-service external to the [[Adrestia-Cluster]]
	* [[Adrestia-fs]]
		* The `adrestia-fs` application is planned as a distributed, POSIX-compliant file system which stores blocks as objects in the system.
		* This tool will exist as both a client and service external to the [[Adrestia-Cluster]]
	* [[Adrestia-GraphDb]]
		* The `adrestia-graphdb` service will run external to the [[Adrestia-Cluster]], providing a Graph Database where edges and vectors in the graphs are stored as objects in the cluster.
		* The `adrestia-graphdb` should use cryptographic means to segment individual graph database instances from one another such that two instances of the service can only see their respective graphs but not the graphs of one another.
* Design Parameters
	* The Client Application interacts with the [[Adrestia-Daemon]] through the [[Adrestia-Common]] library only.
	* The Client Application maintains zero awareness of the [[Adrestia-Cluster]] internals, including what objects are stored and where they are stored.
	* The integrity of the Client Application and [[Adrestia-Daemon]] are preserved only  through code-signing strategies.  
	* Secure connectivity between Client Application and [[Adrestia-Daemon]] is left to the local operating system.
