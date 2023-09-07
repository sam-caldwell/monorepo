package args

// generalOptions - Evaluate whether general options exist in command line
func (arg *Arguments) generalOptions(args []string) (remainingArgs []string) {
	for len(args) > 0 {
		switch args[0] {
		case "--debug":
			args = removeElement0(args)
			arg.option.debug = true
		case "--color":
			args = removeElement0(args)
			arg.option.color = true
		case "--noop":
			args = removeElement0(args)
			arg.option.noop = true
		default:
			return args
		}
	}
	return args
}

func removeElement0(slice []string) []string {
	return slice[1:]
}
