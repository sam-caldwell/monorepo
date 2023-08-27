package argparse

// Postscript - Add another line to the postfix set
func (arg *Arguments) Postscript(line string) *Arguments {
	_ = arg.err.Push(arg.postscriptText.Add(line))
	return arg
}
