package runcommand

// MockCommandExecutor - A mock implementation of the CommandExecutor
type MockCommandExecutor struct {
	ProcessFunc func(m *MockCommandExecutor, name string, arg ...string) ([]byte, error)
	Outputs     map[string]any
	Errors      map[string]error
}
