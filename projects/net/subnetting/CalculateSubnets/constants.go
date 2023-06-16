package CalculateSubnets

const (
	//ArgParentCIDR - cli Argument position
	ArgParentCIDR = 1

	//ArgSubnetSize - cli Argument position
	ArgSubnetSize = 2

	//ArgResultCount - cli Argument position
	ArgResultCount = 3

	// ExitSuccess - standard exit code
	ExitSuccess = 0

	// ExitMissingArgs  - standard exit code
	ExitMissingArgs = 1

	// ExitSubnettingError  - standard exit code
	ExitSubnettingError = 2

	//ExitInvalidResultCount - standard exit code
	ExitInvalidResultCount = 3

	// ErrGeneral - General error formatting string
	ErrGeneral = "Error:%s"

	//ErrMissingArguments - standard error
	ErrMissingArguments = "Error: missing arguments parentCIDR subnetSize"

	// ErrInvalidParentCIDR - standard error
	ErrInvalidParentCIDR = "invalid parent CIDR:%s"

	// ErrInvalidSubnetSize - standard error
	ErrInvalidSubnetSize = "invalid subnet size:%d"

	// ErrInvalidResultCount - standard error
	ErrInvalidResultCount = "invalid (optional) result count (expect number >0)"

	//MsgIpv4CIDR - standard CIDR format string
	MsgIpv4CIDR = "%d.%d.%d.%d/%d\n"
)
