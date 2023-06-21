package runcommand

func (m MockCommandExecutor) Execute(name string, arg ...string) ([]byte, error) {
	return []byte(m.Output), m.Error
}
