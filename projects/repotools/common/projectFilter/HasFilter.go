package projectFilter

// HasFilter - Returns boolean if a given filter bit is set
func (o *Filter) HasFilter(filter Filter) bool {
	return (*o & filter) == filter
}
