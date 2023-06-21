package runcommand

// Execute - Intercept a command execution and instead return predefined response content
func (m MockCommandExecutor) Execute(name string, arg ...string) ([]byte, error) {
	return []byte(m.Output), m.Error
}
