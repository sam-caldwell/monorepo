package descriptormap

// Count - return a count of argument descriptors
func (m *Map) Count() int {
	if m.data == nil {
		return 0
	}
	return len(m.data)
}
