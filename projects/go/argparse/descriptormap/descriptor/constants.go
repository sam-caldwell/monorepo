package descriptor

const (
	argHelpLong  = "--help"
	argHelpName  = "help"
	argHelpShort = "-h"
	hyphen       = "-"
)

const (
	errArgExists                      = "duplicate argument encountered"
	errReservedArg                    = "commandline descriptor %s is reserved"
	errPositionalArgumentCannotBeFlag = "a positional argument cannot be a flag"
)
