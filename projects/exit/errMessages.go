package exit

const (
	ErrDuplicateEntry           = "duplicate entry"
	ErrMissingArguments         = "missing argument"
	ErrMissingArgWithDetail     = "missing argument (%s)"
	ErrIndexOutOfRange          = "index out of range"
	ErrInternalError            = "internal error"
	ErrInvalidCommand           = "unknown command"
	ErrInvalidCommandWithDetail = "unknown command (%s)"
	ErrInvalidContextId         = "invalid contextId (%s)"
	ErrInvalidInput             = "invalid input"
	ErrLockCheckFailed          = "lock check failed"
	ErrMissingColor             = "Missing color"
	ErrMissingContextId         = "missing contextId (uuid)"
	ErrNotFound                 = "not found"
	ErrNotInitialized           = "not initialized"
	ErrTypeMismatch             = "type mismatch"
	ErrUnknownCommand           = "Unknown command"

	ErrUnsupportedOpsys           = "unsupported operating system"
	ErrUnsupportedOpsysWithDetail = "unsupported operating system: %s"
)
