package hijack

// delete - delete a record from our TrackerTable map
func (table *TrackerTable) delete(key uintptr) (err error) {
	delete(table.patches, key)
	return nil
}
