package runcommand

// SimpleMockCommandExecutor - A mock implementation of the CommandExecutor
type SimpleMockCommandExecutor struct {
	ProcessFunc func(m *SimpleMockCommandExecutor, name string, arg ...string) ([]byte, error)
	Outputs     map[string]any
	Errors      map[string]error
}
