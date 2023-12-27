1. Each Availability Zone (AZ) must have a public key infrastructure (PKI) service instance.
2. The PKI instance will--
	1. Issue a Certificate Authority (CA) for the cluster.
	2. Issue an Intermediate CA (ICA) for each Availability Zone.
	3. Issue certificates signed by the ICA of each Availability Zone (AZ) for each node in the cluster.
	4. Issue Client Certificates to any Adrestia Daemon instance registering itself with an Edge node.
	5. Publish/expose a certificate revocation list (CRL) used to evict nodes/clients from the cluster.
	6. Expose an API for requesting, signing and revoking PKI certificates.
3. During cluster provisioning, the PKI instance for an AZ must be the first instance deployed, and during deployment this PKI instance will establish the "trust root" (CA) and ICA for the cluster.