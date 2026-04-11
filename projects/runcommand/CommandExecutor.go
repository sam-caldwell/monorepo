package runcommand

// CommandExecutor - An interface for both real and mock command execution.
type CommandExecutor interface {
	Execute(name string, arg ...string) ([]byte, error)
}
