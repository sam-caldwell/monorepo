package descriptormap

// CountPositionalArgs - return a count of positional arguments
func (m *Map) CountPositionalArgs() (count int) {
	for _, argument := range m.data {
		if (argument.GetShort() == "") || (argument.GetLong() == "") {
			count++
		}
	}
	return count
}
