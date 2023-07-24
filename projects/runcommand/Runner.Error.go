package runcommand

// Error - command-runner terminating node
func (run *Runner) Error() error {
	return run.err
}
