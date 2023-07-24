package runcommand

// Run - Command runner executor node.
func (run *Runner) Run(command string) *Runner {
	if run.err != nil {
		return run
	}
	run.out, run.err = ShellExecute(command)
	return run
}

// Run - Create a new Runner and execute a command.  This starts a chain.
func Run(command string) *Runner {
	var r Runner
	return r.Run(command)
}
