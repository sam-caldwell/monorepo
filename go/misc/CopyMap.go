package misc

// CopyMap is a helper function to create a copy of the map.
func CopyMap[KeyType comparable, ValueType any](dest, src map[KeyType]ValueType) {
	PurgeMap(&dest) //Delete any contents from the destination to avoid merging two maps.
	for key, value := range src {
		dest[key] = value
	}
}
