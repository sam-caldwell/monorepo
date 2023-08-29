package hijack

// NewTrackerTable - Create, initialize and return a reference to our TrackerTable
func NewTrackerTable() (table *TrackerTable) {
	table = &(TrackerTable{})
	table.lock.Lock()
	defer table.lock.Unlock()
	table.patches = make(map[uintptr]AppliedPatch)
	return table
}
