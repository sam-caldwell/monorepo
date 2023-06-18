package exit

const (
	GeneralError   = 1
	InvalidCommand = 2
	InvalidInput   = 3

	MissingArg = 10

	NotFound = 20

	LockCreateFailed = 98
	FailedToFreeLock = 99

	ErrMissingArgCommand = "missing argument (%s)"
	ErrInvalidCommand    = "unknown command (%s)"
	ErrMissingContextId  = "missing contextId (uuid)"
	ErrInvalidContextId  = "invalid contextId (%s)"
	ErrLockCheckFailed   = "lock check failed"

	errMessage = "Error: %s%s"
	errUsage   = "\n\nUsage:\n%s\n"
)
