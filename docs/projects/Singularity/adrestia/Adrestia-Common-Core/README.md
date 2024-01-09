
# Objectives
* Implement a framework for building applications within the Adrestia cluster and beyond.
* Create a framework which will create distributed applications without regard for underlying platform, as autonomous nodes capable of self-discovery of their peer instances.
* Ensure zero trust in an ecosystem of heterogenous systems by building self-contained services with zero platform dependencies.

## Adrestia Abstraction Model![[common-core-drawing]]
## [[Router Layer]]
The router layer maintains connectivity between clients and a dynamically provisioned fleet of service instances, exposing the service layer endpoints through virtual IP addresses (VIPs) and routing traffic to the nearest possible service instances.

Router Layer Features:
* Dynamic routing through virtual IP address facilities.
* Rate Limiting and Load Shedding.
* Instance health checks.
## Security Layer
The security layer facilitates the public-key infrastructure (PKI) and zero-trust model between clients and service instances as well as between service-peer instances themselves.  The security layer also facilitates encryption at-rest by providing data read/write encryption interfaces and keys, replicated between the homogenous peers of a given service instance type.

Security Layer Features:
* Cluster-specific Certificate Authority (CA).
* Certificate signing features for peers.
* Mutual cryptographic authentication.
* Trust revocation.
* Encryption in-flight tunneling between peers.
* Encryption in-flight between clients and instances.
* Encryption at-rest facilities for read/writes to persistent storage.
## Config Layer
The configuration layer is the reusable component of the framework which allows the service layer to store replicated configuration map data between homogenous instances.  For example, a service A which produces some given value to the end user could exist as one or more instances with a shared configuration map readable and writable by the service layer, replicated between its peers.

Configuration Layer Features
* Distributed in-memory key-store
* Replication of key-store across homogenous cluster instances.
* Key-Store integrity validation.
## Service Layer
The service layer is the application tier where most programming is performed.  This is the layer where application-specific code leverages the other layers to deliver end-user value.