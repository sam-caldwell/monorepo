package runcommand

// Run - Command runner executor node.
func (run *Runner) Run(command string) *Runner {
	if run.err != nil {
		return run
	}
	var stdout string
	stdout, run.err = ShellExecute(command)
	run.out = run.out + "{command:'" + command + "',\n" +
		"output:'" + stdout + "'}\n"
	return run
}

// Run - Create a new Runner and execute a command.  This starts a chain.
func Run(command string) *Runner {
	var r Runner
	return r.Run(command)
}
