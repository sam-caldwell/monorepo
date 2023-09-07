package args

// generalOptions - Evaluate whether general options exist in command line
func (arg *Arguments) generalOptions(args []string) (remainingArgs []string) {
	deleteFirstElement := func(slice []string) []string {
		return slice[1:]
	}
	for len(args) > 0 {
		switch args[0] {
		case "--debug":
			args = deleteFirstElement(args)
			arg.option.debug = true
		case "--color":
			args = deleteFirstElement(args)
			arg.option.color = true
		case "--noop":
			args = deleteFirstElement(args)
			arg.option.noop = true
		default:
			return args
		}
	}
	return args
}
