package argparse

// Dump - Dump the descriptor names as a list of strings and their values
func (arg *Arguments) Dump() (output []string) {
	for key := range arg.descriptors.List() {
		output = append(output, key)
	}
	return output
}
