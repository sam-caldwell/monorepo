package argparse

const (
	// Exit codes
	exitArgParseError = 1

	// Error messages
	//errParsingCliArgs = "error parsing command-line arguments at %s"

	errArgParseError = "Error parsing command-line arguments"

	errFoundUndefinedArgument = "found undefined argument '%s'"

	errPositionalArgumentFollowingOptional = "error: positional arguments cannot follow optional arguments (position: %d)"
)
