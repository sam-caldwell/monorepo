package runcommand

type CommandExecutor interface {
	Execute(name string, arg ...string) ([]byte, error)
}
