package monorepo

// commandArgs - parse the commandline for command and arguments
func commandArgs(commandLine []string) (command string, args []string) {
	if len(commandLine[1:]) > 0 {
		args = commandLine[1:]
	}
	return commandLine[0], args
}
