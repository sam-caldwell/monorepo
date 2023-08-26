package argparse

// Count - Return a descriptor count
func (arg *Arguments) Count() int {
	return (*arg).descriptors.Count()
}
