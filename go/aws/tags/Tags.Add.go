package tags

// Add - Add a new element to the map of key-value pairs
func (tags *Tags) Add(key, value string) {
	if tags.data == nil {
		tags.data = make(map[string]string)
	}
	//ToDo: validate the key
	tags.data[key] = value
}
