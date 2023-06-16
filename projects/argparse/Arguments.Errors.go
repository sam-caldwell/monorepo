package argparse

// Errors return a list of errors
func (arg *Arguments) Errors() (result []string) {
	for _, e := range arg.err.List() {
		result = append(result, e.(error).Error())
	}
	return result
}
