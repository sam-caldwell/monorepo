* 5 Jan 2024 (v0.0.6)
	* `databases/sql/tracker`: Add `isValidEmail()` SQL function.
	* `go/pkicore`: finishing `CreateCa()` and unit test. This function creates a Certificate Authority (CA) root certificate key pair.
	* `go/pkicore`: finishing `CreateCertificate()` and unit test. This function creates a certificate key-pair (intermediate CA or end-user CA).
	* `go/tools/monorepo`: Fix bug.  Forgot to take move the dependency-order sorted list to the output.
	* `containers/opsys`: Fix missing tag statement for ubuntu:latest which was masked by not building from a clean state previously.  Build now works on clean systems.
	
* 31 Dec 2023 (v0.0.5)
	* `go/tools/monorepo`: Added dependency sorting so our build order is not lexically sorted.  This fixes a bug discovered when building from scratch on freshly provisioned system.