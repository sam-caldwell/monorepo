package projectFilter

// String - return string value for the given filter.
func (o *Filter) String() (result []string) {
	flags := FlagList()
	for _, flag := range flags {
		if o.HasFilter(flag) {
			s, _ := numberToString(flag)
			result = append(result, s)
		}
	}
	return result
}
