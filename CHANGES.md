* 1 Jan 2024
	* `go/tools/monorepo`: Fix bug.  Forgot to take move the dependency-order sorted list to the output.
	* `containers/opsys`: Fix missing tag statement for ubuntu:latest which was masked by not building from a clean state previously.  Build now works on clean systems.
	* `go/pkicore`: finishing `GenerateSelfSigned()` and unit test.
	
* 31 Dec 2023 (v.0.0.5)
	* `go/tools/monorepo`: Added dependency sorting so our build order is not lexically sorted.  This fixes a bug discovered when building from scratch on freshly provisioned system.