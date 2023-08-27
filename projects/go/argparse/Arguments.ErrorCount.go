package argparse

// ErrorCount - return a count of errors
func (arg *Arguments) ErrorCount() int {
	return arg.err.Count()
}
