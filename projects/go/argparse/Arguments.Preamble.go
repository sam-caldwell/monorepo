package argparse

// Preamble - Add another line to the prefix set
func (arg *Arguments) Preamble(line string) *Arguments {
	_ = arg.err.Push(arg.preambleText.Add(line))
	return arg
}
